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
	"github.com/aws/aws-sdk-go/service/auditmanager"
	"github.com/aws/aws-sdk-go/service/auditmanager/auditmanageriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AuditManager struct {
	auditmanageriface.AuditManagerAPI
	cache *cache.Cache
}

func NewAuditManager(auditmanagerapi auditmanageriface.AuditManagerAPI) *AuditManager {
	return &AuditManager{
		AuditManagerAPI: auditmanagerapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AuditManager) BatchAssociateAssessmentReportEvidence(input *auditmanager.BatchAssociateAssessmentReportEvidenceInput) (*auditmanager.BatchAssociateAssessmentReportEvidenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchAssociateAssessmentReportEvidenceOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchAssociateAssessmentReportEvidence(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchAssociateAssessmentReportEvidenceWithContext(ctx context.Context, input *auditmanager.BatchAssociateAssessmentReportEvidenceInput, opts ...request.Option) (*auditmanager.BatchAssociateAssessmentReportEvidenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchAssociateAssessmentReportEvidenceOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchAssociateAssessmentReportEvidenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchCreateDelegationByAssessment(input *auditmanager.BatchCreateDelegationByAssessmentInput) (*auditmanager.BatchCreateDelegationByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchCreateDelegationByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchCreateDelegationByAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchCreateDelegationByAssessmentWithContext(ctx context.Context, input *auditmanager.BatchCreateDelegationByAssessmentInput, opts ...request.Option) (*auditmanager.BatchCreateDelegationByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchCreateDelegationByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchCreateDelegationByAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchDeleteDelegationByAssessment(input *auditmanager.BatchDeleteDelegationByAssessmentInput) (*auditmanager.BatchDeleteDelegationByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchDeleteDelegationByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchDeleteDelegationByAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchDeleteDelegationByAssessmentWithContext(ctx context.Context, input *auditmanager.BatchDeleteDelegationByAssessmentInput, opts ...request.Option) (*auditmanager.BatchDeleteDelegationByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchDeleteDelegationByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchDeleteDelegationByAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchDisassociateAssessmentReportEvidence(input *auditmanager.BatchDisassociateAssessmentReportEvidenceInput) (*auditmanager.BatchDisassociateAssessmentReportEvidenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchDisassociateAssessmentReportEvidenceOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchDisassociateAssessmentReportEvidence(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchDisassociateAssessmentReportEvidenceWithContext(ctx context.Context, input *auditmanager.BatchDisassociateAssessmentReportEvidenceInput, opts ...request.Option) (*auditmanager.BatchDisassociateAssessmentReportEvidenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchDisassociateAssessmentReportEvidenceOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchDisassociateAssessmentReportEvidenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchImportEvidenceToAssessmentControl(input *auditmanager.BatchImportEvidenceToAssessmentControlInput) (*auditmanager.BatchImportEvidenceToAssessmentControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchImportEvidenceToAssessmentControlOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchImportEvidenceToAssessmentControl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) BatchImportEvidenceToAssessmentControlWithContext(ctx context.Context, input *auditmanager.BatchImportEvidenceToAssessmentControlInput, opts ...request.Option) (*auditmanager.BatchImportEvidenceToAssessmentControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.BatchImportEvidenceToAssessmentControlOutput), nil
	}
	out, err := c.AuditManagerAPI.BatchImportEvidenceToAssessmentControlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAccountStatus(input *auditmanager.GetAccountStatusInput) (*auditmanager.GetAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAccountStatusOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAccountStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAccountStatusWithContext(ctx context.Context, input *auditmanager.GetAccountStatusInput, opts ...request.Option) (*auditmanager.GetAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAccountStatusOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAccountStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAssessment(input *auditmanager.GetAssessmentInput) (*auditmanager.GetAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAssessmentFramework(input *auditmanager.GetAssessmentFrameworkInput) (*auditmanager.GetAssessmentFrameworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAssessmentFrameworkOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAssessmentFramework(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAssessmentFrameworkWithContext(ctx context.Context, input *auditmanager.GetAssessmentFrameworkInput, opts ...request.Option) (*auditmanager.GetAssessmentFrameworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAssessmentFrameworkOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAssessmentFrameworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAssessmentReportUrl(input *auditmanager.GetAssessmentReportUrlInput) (*auditmanager.GetAssessmentReportUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAssessmentReportUrlOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAssessmentReportUrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAssessmentReportUrlWithContext(ctx context.Context, input *auditmanager.GetAssessmentReportUrlInput, opts ...request.Option) (*auditmanager.GetAssessmentReportUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAssessmentReportUrlOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAssessmentReportUrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetAssessmentWithContext(ctx context.Context, input *auditmanager.GetAssessmentInput, opts ...request.Option) (*auditmanager.GetAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.GetAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetChangeLogs(input *auditmanager.GetChangeLogsInput) (*auditmanager.GetChangeLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetChangeLogsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetChangeLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetChangeLogsPages(input *auditmanager.GetChangeLogsInput, fn func(*auditmanager.GetChangeLogsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetChangeLogsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetChangeLogsOutput{}
	fnCacher := func(out *auditmanager.GetChangeLogsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetChangeLogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetChangeLogsPagesWithContext(ctx context.Context, input *auditmanager.GetChangeLogsInput, fn func(*auditmanager.GetChangeLogsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetChangeLogsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetChangeLogsOutput{}
	fnCacher := func(out *auditmanager.GetChangeLogsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetChangeLogsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetChangeLogsWithContext(ctx context.Context, input *auditmanager.GetChangeLogsInput, opts ...request.Option) (*auditmanager.GetChangeLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetChangeLogsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetChangeLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetControl(input *auditmanager.GetControlInput) (*auditmanager.GetControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetControlOutput), nil
	}
	out, err := c.AuditManagerAPI.GetControl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetControlWithContext(ctx context.Context, input *auditmanager.GetControlInput, opts ...request.Option) (*auditmanager.GetControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetControlOutput), nil
	}
	out, err := c.AuditManagerAPI.GetControlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetDelegations(input *auditmanager.GetDelegationsInput) (*auditmanager.GetDelegationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetDelegationsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetDelegations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetDelegationsPages(input *auditmanager.GetDelegationsInput, fn func(*auditmanager.GetDelegationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetDelegationsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetDelegationsOutput{}
	fnCacher := func(out *auditmanager.GetDelegationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetDelegationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetDelegationsPagesWithContext(ctx context.Context, input *auditmanager.GetDelegationsInput, fn func(*auditmanager.GetDelegationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetDelegationsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetDelegationsOutput{}
	fnCacher := func(out *auditmanager.GetDelegationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetDelegationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetDelegationsWithContext(ctx context.Context, input *auditmanager.GetDelegationsInput, opts ...request.Option) (*auditmanager.GetDelegationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetDelegationsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetDelegationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidence(input *auditmanager.GetEvidenceInput) (*auditmanager.GetEvidenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidence(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceByEvidenceFolder(input *auditmanager.GetEvidenceByEvidenceFolderInput) (*auditmanager.GetEvidenceByEvidenceFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceByEvidenceFolderOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceByEvidenceFolder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceByEvidenceFolderPages(input *auditmanager.GetEvidenceByEvidenceFolderInput, fn func(*auditmanager.GetEvidenceByEvidenceFolderOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetEvidenceByEvidenceFolderOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetEvidenceByEvidenceFolderOutput{}
	fnCacher := func(out *auditmanager.GetEvidenceByEvidenceFolderOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetEvidenceByEvidenceFolderPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetEvidenceByEvidenceFolderPagesWithContext(ctx context.Context, input *auditmanager.GetEvidenceByEvidenceFolderInput, fn func(*auditmanager.GetEvidenceByEvidenceFolderOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetEvidenceByEvidenceFolderOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetEvidenceByEvidenceFolderOutput{}
	fnCacher := func(out *auditmanager.GetEvidenceByEvidenceFolderOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetEvidenceByEvidenceFolderPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetEvidenceByEvidenceFolderWithContext(ctx context.Context, input *auditmanager.GetEvidenceByEvidenceFolderInput, opts ...request.Option) (*auditmanager.GetEvidenceByEvidenceFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceByEvidenceFolderOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceByEvidenceFolderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceFolder(input *auditmanager.GetEvidenceFolderInput) (*auditmanager.GetEvidenceFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceFolderOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceFolder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceFolderWithContext(ctx context.Context, input *auditmanager.GetEvidenceFolderInput, opts ...request.Option) (*auditmanager.GetEvidenceFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceFolderOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceFolderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceFoldersByAssessment(input *auditmanager.GetEvidenceFoldersByAssessmentInput) (*auditmanager.GetEvidenceFoldersByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceFoldersByAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentControl(input *auditmanager.GetEvidenceFoldersByAssessmentControlInput) (*auditmanager.GetEvidenceFoldersByAssessmentControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentControlOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentControl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentControlPages(input *auditmanager.GetEvidenceFoldersByAssessmentControlInput, fn func(*auditmanager.GetEvidenceFoldersByAssessmentControlOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentControlOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetEvidenceFoldersByAssessmentControlOutput{}
	fnCacher := func(out *auditmanager.GetEvidenceFoldersByAssessmentControlOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentControlPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentControlPagesWithContext(ctx context.Context, input *auditmanager.GetEvidenceFoldersByAssessmentControlInput, fn func(*auditmanager.GetEvidenceFoldersByAssessmentControlOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentControlOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetEvidenceFoldersByAssessmentControlOutput{}
	fnCacher := func(out *auditmanager.GetEvidenceFoldersByAssessmentControlOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentControlPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentControlWithContext(ctx context.Context, input *auditmanager.GetEvidenceFoldersByAssessmentControlInput, opts ...request.Option) (*auditmanager.GetEvidenceFoldersByAssessmentControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentControlOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentControlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentPages(input *auditmanager.GetEvidenceFoldersByAssessmentInput, fn func(*auditmanager.GetEvidenceFoldersByAssessmentOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetEvidenceFoldersByAssessmentOutput{}
	fnCacher := func(out *auditmanager.GetEvidenceFoldersByAssessmentOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentPagesWithContext(ctx context.Context, input *auditmanager.GetEvidenceFoldersByAssessmentInput, fn func(*auditmanager.GetEvidenceFoldersByAssessmentOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.GetEvidenceFoldersByAssessmentOutput{}
	fnCacher := func(out *auditmanager.GetEvidenceFoldersByAssessmentOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) GetEvidenceFoldersByAssessmentWithContext(ctx context.Context, input *auditmanager.GetEvidenceFoldersByAssessmentInput, opts ...request.Option) (*auditmanager.GetEvidenceFoldersByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceFoldersByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceFoldersByAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetEvidenceWithContext(ctx context.Context, input *auditmanager.GetEvidenceInput, opts ...request.Option) (*auditmanager.GetEvidenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetEvidenceOutput), nil
	}
	out, err := c.AuditManagerAPI.GetEvidenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetInsights(input *auditmanager.GetInsightsInput) (*auditmanager.GetInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetInsightsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetInsightsByAssessment(input *auditmanager.GetInsightsByAssessmentInput) (*auditmanager.GetInsightsByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetInsightsByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.GetInsightsByAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetInsightsByAssessmentWithContext(ctx context.Context, input *auditmanager.GetInsightsByAssessmentInput, opts ...request.Option) (*auditmanager.GetInsightsByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetInsightsByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.GetInsightsByAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetInsightsWithContext(ctx context.Context, input *auditmanager.GetInsightsInput, opts ...request.Option) (*auditmanager.GetInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetInsightsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetOrganizationAdminAccount(input *auditmanager.GetOrganizationAdminAccountInput) (*auditmanager.GetOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetOrganizationAdminAccountOutput), nil
	}
	out, err := c.AuditManagerAPI.GetOrganizationAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetOrganizationAdminAccountWithContext(ctx context.Context, input *auditmanager.GetOrganizationAdminAccountInput, opts ...request.Option) (*auditmanager.GetOrganizationAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetOrganizationAdminAccountOutput), nil
	}
	out, err := c.AuditManagerAPI.GetOrganizationAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetServicesInScope(input *auditmanager.GetServicesInScopeInput) (*auditmanager.GetServicesInScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetServicesInScopeOutput), nil
	}
	out, err := c.AuditManagerAPI.GetServicesInScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetServicesInScopeWithContext(ctx context.Context, input *auditmanager.GetServicesInScopeInput, opts ...request.Option) (*auditmanager.GetServicesInScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetServicesInScopeOutput), nil
	}
	out, err := c.AuditManagerAPI.GetServicesInScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetSettings(input *auditmanager.GetSettingsInput) (*auditmanager.GetSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetSettingsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) GetSettingsWithContext(ctx context.Context, input *auditmanager.GetSettingsInput, opts ...request.Option) (*auditmanager.GetSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.GetSettingsOutput), nil
	}
	out, err := c.AuditManagerAPI.GetSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentControlInsightsByControlDomain(input *auditmanager.ListAssessmentControlInsightsByControlDomainInput) (*auditmanager.ListAssessmentControlInsightsByControlDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentControlInsightsByControlDomainOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentControlInsightsByControlDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentControlInsightsByControlDomainPages(input *auditmanager.ListAssessmentControlInsightsByControlDomainInput, fn func(*auditmanager.ListAssessmentControlInsightsByControlDomainOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentControlInsightsByControlDomainOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentControlInsightsByControlDomainOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentControlInsightsByControlDomainOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentControlInsightsByControlDomainPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentControlInsightsByControlDomainPagesWithContext(ctx context.Context, input *auditmanager.ListAssessmentControlInsightsByControlDomainInput, fn func(*auditmanager.ListAssessmentControlInsightsByControlDomainOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentControlInsightsByControlDomainOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentControlInsightsByControlDomainOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentControlInsightsByControlDomainOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentControlInsightsByControlDomainPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentControlInsightsByControlDomainWithContext(ctx context.Context, input *auditmanager.ListAssessmentControlInsightsByControlDomainInput, opts ...request.Option) (*auditmanager.ListAssessmentControlInsightsByControlDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentControlInsightsByControlDomainOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentControlInsightsByControlDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentFrameworkShareRequests(input *auditmanager.ListAssessmentFrameworkShareRequestsInput) (*auditmanager.ListAssessmentFrameworkShareRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentFrameworkShareRequestsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentFrameworkShareRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentFrameworkShareRequestsPages(input *auditmanager.ListAssessmentFrameworkShareRequestsInput, fn func(*auditmanager.ListAssessmentFrameworkShareRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentFrameworkShareRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentFrameworkShareRequestsOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentFrameworkShareRequestsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentFrameworkShareRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentFrameworkShareRequestsPagesWithContext(ctx context.Context, input *auditmanager.ListAssessmentFrameworkShareRequestsInput, fn func(*auditmanager.ListAssessmentFrameworkShareRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentFrameworkShareRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentFrameworkShareRequestsOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentFrameworkShareRequestsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentFrameworkShareRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentFrameworkShareRequestsWithContext(ctx context.Context, input *auditmanager.ListAssessmentFrameworkShareRequestsInput, opts ...request.Option) (*auditmanager.ListAssessmentFrameworkShareRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentFrameworkShareRequestsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentFrameworkShareRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentFrameworks(input *auditmanager.ListAssessmentFrameworksInput) (*auditmanager.ListAssessmentFrameworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentFrameworksOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentFrameworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentFrameworksPages(input *auditmanager.ListAssessmentFrameworksInput, fn func(*auditmanager.ListAssessmentFrameworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentFrameworksOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentFrameworksOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentFrameworksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentFrameworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentFrameworksPagesWithContext(ctx context.Context, input *auditmanager.ListAssessmentFrameworksInput, fn func(*auditmanager.ListAssessmentFrameworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentFrameworksOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentFrameworksOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentFrameworksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentFrameworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentFrameworksWithContext(ctx context.Context, input *auditmanager.ListAssessmentFrameworksInput, opts ...request.Option) (*auditmanager.ListAssessmentFrameworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentFrameworksOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentFrameworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentReports(input *auditmanager.ListAssessmentReportsInput) (*auditmanager.ListAssessmentReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentReportsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentReportsPages(input *auditmanager.ListAssessmentReportsInput, fn func(*auditmanager.ListAssessmentReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentReportsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentReportsOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentReportsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentReportsPagesWithContext(ctx context.Context, input *auditmanager.ListAssessmentReportsInput, fn func(*auditmanager.ListAssessmentReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentReportsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentReportsOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentReportsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentReportsWithContext(ctx context.Context, input *auditmanager.ListAssessmentReportsInput, opts ...request.Option) (*auditmanager.ListAssessmentReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentReportsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessments(input *auditmanager.ListAssessmentsInput) (*auditmanager.ListAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListAssessmentsPages(input *auditmanager.ListAssessmentsInput, fn func(*auditmanager.ListAssessmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentsOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentsPagesWithContext(ctx context.Context, input *auditmanager.ListAssessmentsInput, fn func(*auditmanager.ListAssessmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListAssessmentsOutput{}
	fnCacher := func(out *auditmanager.ListAssessmentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListAssessmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListAssessmentsWithContext(ctx context.Context, input *auditmanager.ListAssessmentsInput, opts ...request.Option) (*auditmanager.ListAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListAssessmentsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListAssessmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlDomainInsights(input *auditmanager.ListControlDomainInsightsInput) (*auditmanager.ListControlDomainInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlDomainInsightsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlDomainInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlDomainInsightsByAssessment(input *auditmanager.ListControlDomainInsightsByAssessmentInput) (*auditmanager.ListControlDomainInsightsByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlDomainInsightsByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlDomainInsightsByAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlDomainInsightsByAssessmentPages(input *auditmanager.ListControlDomainInsightsByAssessmentInput, fn func(*auditmanager.ListControlDomainInsightsByAssessmentOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlDomainInsightsByAssessmentOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlDomainInsightsByAssessmentOutput{}
	fnCacher := func(out *auditmanager.ListControlDomainInsightsByAssessmentOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlDomainInsightsByAssessmentPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlDomainInsightsByAssessmentPagesWithContext(ctx context.Context, input *auditmanager.ListControlDomainInsightsByAssessmentInput, fn func(*auditmanager.ListControlDomainInsightsByAssessmentOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlDomainInsightsByAssessmentOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlDomainInsightsByAssessmentOutput{}
	fnCacher := func(out *auditmanager.ListControlDomainInsightsByAssessmentOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlDomainInsightsByAssessmentPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlDomainInsightsByAssessmentWithContext(ctx context.Context, input *auditmanager.ListControlDomainInsightsByAssessmentInput, opts ...request.Option) (*auditmanager.ListControlDomainInsightsByAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlDomainInsightsByAssessmentOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlDomainInsightsByAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlDomainInsightsPages(input *auditmanager.ListControlDomainInsightsInput, fn func(*auditmanager.ListControlDomainInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlDomainInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlDomainInsightsOutput{}
	fnCacher := func(out *auditmanager.ListControlDomainInsightsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlDomainInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlDomainInsightsPagesWithContext(ctx context.Context, input *auditmanager.ListControlDomainInsightsInput, fn func(*auditmanager.ListControlDomainInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlDomainInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlDomainInsightsOutput{}
	fnCacher := func(out *auditmanager.ListControlDomainInsightsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlDomainInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlDomainInsightsWithContext(ctx context.Context, input *auditmanager.ListControlDomainInsightsInput, opts ...request.Option) (*auditmanager.ListControlDomainInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlDomainInsightsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlDomainInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlInsightsByControlDomain(input *auditmanager.ListControlInsightsByControlDomainInput) (*auditmanager.ListControlInsightsByControlDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlInsightsByControlDomainOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlInsightsByControlDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlInsightsByControlDomainPages(input *auditmanager.ListControlInsightsByControlDomainInput, fn func(*auditmanager.ListControlInsightsByControlDomainOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlInsightsByControlDomainOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlInsightsByControlDomainOutput{}
	fnCacher := func(out *auditmanager.ListControlInsightsByControlDomainOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlInsightsByControlDomainPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlInsightsByControlDomainPagesWithContext(ctx context.Context, input *auditmanager.ListControlInsightsByControlDomainInput, fn func(*auditmanager.ListControlInsightsByControlDomainOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlInsightsByControlDomainOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlInsightsByControlDomainOutput{}
	fnCacher := func(out *auditmanager.ListControlInsightsByControlDomainOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlInsightsByControlDomainPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlInsightsByControlDomainWithContext(ctx context.Context, input *auditmanager.ListControlInsightsByControlDomainInput, opts ...request.Option) (*auditmanager.ListControlInsightsByControlDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlInsightsByControlDomainOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlInsightsByControlDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControls(input *auditmanager.ListControlsInput) (*auditmanager.ListControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListControlsPages(input *auditmanager.ListControlsInput, fn func(*auditmanager.ListControlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlsOutput{}
	fnCacher := func(out *auditmanager.ListControlsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlsPagesWithContext(ctx context.Context, input *auditmanager.ListControlsInput, fn func(*auditmanager.ListControlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListControlsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListControlsOutput{}
	fnCacher := func(out *auditmanager.ListControlsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListControlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListControlsWithContext(ctx context.Context, input *auditmanager.ListControlsInput, opts ...request.Option) (*auditmanager.ListControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListControlsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListKeywordsForDataSource(input *auditmanager.ListKeywordsForDataSourceInput) (*auditmanager.ListKeywordsForDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListKeywordsForDataSourceOutput), nil
	}
	out, err := c.AuditManagerAPI.ListKeywordsForDataSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListKeywordsForDataSourcePages(input *auditmanager.ListKeywordsForDataSourceInput, fn func(*auditmanager.ListKeywordsForDataSourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListKeywordsForDataSourceOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListKeywordsForDataSourceOutput{}
	fnCacher := func(out *auditmanager.ListKeywordsForDataSourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListKeywordsForDataSourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListKeywordsForDataSourcePagesWithContext(ctx context.Context, input *auditmanager.ListKeywordsForDataSourceInput, fn func(*auditmanager.ListKeywordsForDataSourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListKeywordsForDataSourceOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListKeywordsForDataSourceOutput{}
	fnCacher := func(out *auditmanager.ListKeywordsForDataSourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListKeywordsForDataSourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListKeywordsForDataSourceWithContext(ctx context.Context, input *auditmanager.ListKeywordsForDataSourceInput, opts ...request.Option) (*auditmanager.ListKeywordsForDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListKeywordsForDataSourceOutput), nil
	}
	out, err := c.AuditManagerAPI.ListKeywordsForDataSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListNotifications(input *auditmanager.ListNotificationsInput) (*auditmanager.ListNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListNotificationsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListNotificationsPages(input *auditmanager.ListNotificationsInput, fn func(*auditmanager.ListNotificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListNotificationsOutput{}
	fnCacher := func(out *auditmanager.ListNotificationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListNotificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListNotificationsPagesWithContext(ctx context.Context, input *auditmanager.ListNotificationsInput, fn func(*auditmanager.ListNotificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*auditmanager.ListNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &auditmanager.ListNotificationsOutput{}
	fnCacher := func(out *auditmanager.ListNotificationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AuditManagerAPI.ListNotificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AuditManager) ListNotificationsWithContext(ctx context.Context, input *auditmanager.ListNotificationsInput, opts ...request.Option) (*auditmanager.ListNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListNotificationsOutput), nil
	}
	out, err := c.AuditManagerAPI.ListNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListTagsForResource(input *auditmanager.ListTagsForResourceInput) (*auditmanager.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListTagsForResourceOutput), nil
	}
	out, err := c.AuditManagerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AuditManager) ListTagsForResourceWithContext(ctx context.Context, input *auditmanager.ListTagsForResourceInput, opts ...request.Option) (*auditmanager.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*auditmanager.ListTagsForResourceOutput), nil
	}
	out, err := c.AuditManagerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

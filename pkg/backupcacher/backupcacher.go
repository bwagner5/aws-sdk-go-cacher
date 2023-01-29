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
package backupcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/backup/backupiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Backup struct {
	backupiface.BackupAPI
	cache *cache.Cache
}

func New(backupapi backupiface.BackupAPI) *Backup {
	return &Backup{
		BackupAPI: backupapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Backup) DescribeBackupJob(input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeBackupJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeBackupJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeBackupJobWithContext(ctx context.Context, input *backup.DescribeBackupJobInput, opts ...request.Option) (*backup.DescribeBackupJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeBackupJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeBackupJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeBackupVault(input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeBackupVaultOutput), nil
	}
	out, err := c.BackupAPI.DescribeBackupVault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeBackupVaultWithContext(ctx context.Context, input *backup.DescribeBackupVaultInput, opts ...request.Option) (*backup.DescribeBackupVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeBackupVaultOutput), nil
	}
	out, err := c.BackupAPI.DescribeBackupVaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeCopyJob(input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeCopyJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeCopyJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeCopyJobWithContext(ctx context.Context, input *backup.DescribeCopyJobInput, opts ...request.Option) (*backup.DescribeCopyJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeCopyJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeCopyJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeFramework(input *backup.DescribeFrameworkInput) (*backup.DescribeFrameworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeFrameworkOutput), nil
	}
	out, err := c.BackupAPI.DescribeFramework(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeFrameworkWithContext(ctx context.Context, input *backup.DescribeFrameworkInput, opts ...request.Option) (*backup.DescribeFrameworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeFrameworkOutput), nil
	}
	out, err := c.BackupAPI.DescribeFrameworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeGlobalSettings(input *backup.DescribeGlobalSettingsInput) (*backup.DescribeGlobalSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeGlobalSettingsOutput), nil
	}
	out, err := c.BackupAPI.DescribeGlobalSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeGlobalSettingsWithContext(ctx context.Context, input *backup.DescribeGlobalSettingsInput, opts ...request.Option) (*backup.DescribeGlobalSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeGlobalSettingsOutput), nil
	}
	out, err := c.BackupAPI.DescribeGlobalSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeProtectedResource(input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeProtectedResourceOutput), nil
	}
	out, err := c.BackupAPI.DescribeProtectedResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeProtectedResourceWithContext(ctx context.Context, input *backup.DescribeProtectedResourceInput, opts ...request.Option) (*backup.DescribeProtectedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeProtectedResourceOutput), nil
	}
	out, err := c.BackupAPI.DescribeProtectedResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeRecoveryPoint(input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeRecoveryPointOutput), nil
	}
	out, err := c.BackupAPI.DescribeRecoveryPoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeRecoveryPointWithContext(ctx context.Context, input *backup.DescribeRecoveryPointInput, opts ...request.Option) (*backup.DescribeRecoveryPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeRecoveryPointOutput), nil
	}
	out, err := c.BackupAPI.DescribeRecoveryPointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeRegionSettings(input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeRegionSettingsOutput), nil
	}
	out, err := c.BackupAPI.DescribeRegionSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeRegionSettingsWithContext(ctx context.Context, input *backup.DescribeRegionSettingsInput, opts ...request.Option) (*backup.DescribeRegionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeRegionSettingsOutput), nil
	}
	out, err := c.BackupAPI.DescribeRegionSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeReportJob(input *backup.DescribeReportJobInput) (*backup.DescribeReportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeReportJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeReportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeReportJobWithContext(ctx context.Context, input *backup.DescribeReportJobInput, opts ...request.Option) (*backup.DescribeReportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeReportJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeReportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeReportPlan(input *backup.DescribeReportPlanInput) (*backup.DescribeReportPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeReportPlanOutput), nil
	}
	out, err := c.BackupAPI.DescribeReportPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeReportPlanWithContext(ctx context.Context, input *backup.DescribeReportPlanInput, opts ...request.Option) (*backup.DescribeReportPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeReportPlanOutput), nil
	}
	out, err := c.BackupAPI.DescribeReportPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeRestoreJob(input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeRestoreJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeRestoreJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) DescribeRestoreJobWithContext(ctx context.Context, input *backup.DescribeRestoreJobInput, opts ...request.Option) (*backup.DescribeRestoreJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.DescribeRestoreJobOutput), nil
	}
	out, err := c.BackupAPI.DescribeRestoreJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupPlan(input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupPlanOutput), nil
	}
	out, err := c.BackupAPI.GetBackupPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupPlanFromJSON(input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupPlanFromJSONOutput), nil
	}
	out, err := c.BackupAPI.GetBackupPlanFromJSON(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupPlanFromJSONWithContext(ctx context.Context, input *backup.GetBackupPlanFromJSONInput, opts ...request.Option) (*backup.GetBackupPlanFromJSONOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupPlanFromJSONOutput), nil
	}
	out, err := c.BackupAPI.GetBackupPlanFromJSONWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupPlanFromTemplate(input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupPlanFromTemplateOutput), nil
	}
	out, err := c.BackupAPI.GetBackupPlanFromTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupPlanFromTemplateWithContext(ctx context.Context, input *backup.GetBackupPlanFromTemplateInput, opts ...request.Option) (*backup.GetBackupPlanFromTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupPlanFromTemplateOutput), nil
	}
	out, err := c.BackupAPI.GetBackupPlanFromTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupPlanWithContext(ctx context.Context, input *backup.GetBackupPlanInput, opts ...request.Option) (*backup.GetBackupPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupPlanOutput), nil
	}
	out, err := c.BackupAPI.GetBackupPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupSelection(input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupSelectionOutput), nil
	}
	out, err := c.BackupAPI.GetBackupSelection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupSelectionWithContext(ctx context.Context, input *backup.GetBackupSelectionInput, opts ...request.Option) (*backup.GetBackupSelectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupSelectionOutput), nil
	}
	out, err := c.BackupAPI.GetBackupSelectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupVaultAccessPolicy(input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupVaultAccessPolicyOutput), nil
	}
	out, err := c.BackupAPI.GetBackupVaultAccessPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupVaultAccessPolicyWithContext(ctx context.Context, input *backup.GetBackupVaultAccessPolicyInput, opts ...request.Option) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupVaultAccessPolicyOutput), nil
	}
	out, err := c.BackupAPI.GetBackupVaultAccessPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupVaultNotifications(input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupVaultNotificationsOutput), nil
	}
	out, err := c.BackupAPI.GetBackupVaultNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetBackupVaultNotificationsWithContext(ctx context.Context, input *backup.GetBackupVaultNotificationsInput, opts ...request.Option) (*backup.GetBackupVaultNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetBackupVaultNotificationsOutput), nil
	}
	out, err := c.BackupAPI.GetBackupVaultNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetLegalHold(input *backup.GetLegalHoldInput) (*backup.GetLegalHoldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetLegalHoldOutput), nil
	}
	out, err := c.BackupAPI.GetLegalHold(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetLegalHoldWithContext(ctx context.Context, input *backup.GetLegalHoldInput, opts ...request.Option) (*backup.GetLegalHoldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetLegalHoldOutput), nil
	}
	out, err := c.BackupAPI.GetLegalHoldWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetRecoveryPointRestoreMetadata(input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetRecoveryPointRestoreMetadataOutput), nil
	}
	out, err := c.BackupAPI.GetRecoveryPointRestoreMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetRecoveryPointRestoreMetadataWithContext(ctx context.Context, input *backup.GetRecoveryPointRestoreMetadataInput, opts ...request.Option) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetRecoveryPointRestoreMetadataOutput), nil
	}
	out, err := c.BackupAPI.GetRecoveryPointRestoreMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetSupportedResourceTypes(input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetSupportedResourceTypesOutput), nil
	}
	out, err := c.BackupAPI.GetSupportedResourceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) GetSupportedResourceTypesWithContext(ctx context.Context, input *backup.GetSupportedResourceTypesInput, opts ...request.Option) (*backup.GetSupportedResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.GetSupportedResourceTypesOutput), nil
	}
	out, err := c.BackupAPI.GetSupportedResourceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupJobs(input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupJobsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupJobsPages(input *backup.ListBackupJobsInput, fn func(*backup.ListBackupJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupJobsOutput{}
	fnCacher := func(out *backup.ListBackupJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupJobsPagesWithContext(ctx context.Context, input *backup.ListBackupJobsInput, fn func(*backup.ListBackupJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupJobsOutput{}
	fnCacher := func(out *backup.ListBackupJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupJobsWithContext(ctx context.Context, input *backup.ListBackupJobsInput, opts ...request.Option) (*backup.ListBackupJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupJobsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupPlanTemplates(input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupPlanTemplatesOutput), nil
	}
	out, err := c.BackupAPI.ListBackupPlanTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupPlanTemplatesPages(input *backup.ListBackupPlanTemplatesInput, fn func(*backup.ListBackupPlanTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupPlanTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupPlanTemplatesOutput{}
	fnCacher := func(out *backup.ListBackupPlanTemplatesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupPlanTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupPlanTemplatesPagesWithContext(ctx context.Context, input *backup.ListBackupPlanTemplatesInput, fn func(*backup.ListBackupPlanTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupPlanTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupPlanTemplatesOutput{}
	fnCacher := func(out *backup.ListBackupPlanTemplatesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupPlanTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupPlanTemplatesWithContext(ctx context.Context, input *backup.ListBackupPlanTemplatesInput, opts ...request.Option) (*backup.ListBackupPlanTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupPlanTemplatesOutput), nil
	}
	out, err := c.BackupAPI.ListBackupPlanTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupPlanVersions(input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupPlanVersionsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupPlanVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupPlanVersionsPages(input *backup.ListBackupPlanVersionsInput, fn func(*backup.ListBackupPlanVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupPlanVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupPlanVersionsOutput{}
	fnCacher := func(out *backup.ListBackupPlanVersionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupPlanVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupPlanVersionsPagesWithContext(ctx context.Context, input *backup.ListBackupPlanVersionsInput, fn func(*backup.ListBackupPlanVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupPlanVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupPlanVersionsOutput{}
	fnCacher := func(out *backup.ListBackupPlanVersionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupPlanVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupPlanVersionsWithContext(ctx context.Context, input *backup.ListBackupPlanVersionsInput, opts ...request.Option) (*backup.ListBackupPlanVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupPlanVersionsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupPlanVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupPlans(input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupPlansOutput), nil
	}
	out, err := c.BackupAPI.ListBackupPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupPlansPages(input *backup.ListBackupPlansInput, fn func(*backup.ListBackupPlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupPlansOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupPlansOutput{}
	fnCacher := func(out *backup.ListBackupPlansOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupPlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupPlansPagesWithContext(ctx context.Context, input *backup.ListBackupPlansInput, fn func(*backup.ListBackupPlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupPlansOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupPlansOutput{}
	fnCacher := func(out *backup.ListBackupPlansOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupPlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupPlansWithContext(ctx context.Context, input *backup.ListBackupPlansInput, opts ...request.Option) (*backup.ListBackupPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupPlansOutput), nil
	}
	out, err := c.BackupAPI.ListBackupPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupSelections(input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupSelectionsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupSelections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupSelectionsPages(input *backup.ListBackupSelectionsInput, fn func(*backup.ListBackupSelectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupSelectionsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupSelectionsOutput{}
	fnCacher := func(out *backup.ListBackupSelectionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupSelectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupSelectionsPagesWithContext(ctx context.Context, input *backup.ListBackupSelectionsInput, fn func(*backup.ListBackupSelectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupSelectionsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupSelectionsOutput{}
	fnCacher := func(out *backup.ListBackupSelectionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupSelectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupSelectionsWithContext(ctx context.Context, input *backup.ListBackupSelectionsInput, opts ...request.Option) (*backup.ListBackupSelectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupSelectionsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupSelectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupVaults(input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupVaultsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupVaults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListBackupVaultsPages(input *backup.ListBackupVaultsInput, fn func(*backup.ListBackupVaultsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupVaultsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupVaultsOutput{}
	fnCacher := func(out *backup.ListBackupVaultsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupVaultsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupVaultsPagesWithContext(ctx context.Context, input *backup.ListBackupVaultsInput, fn func(*backup.ListBackupVaultsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListBackupVaultsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListBackupVaultsOutput{}
	fnCacher := func(out *backup.ListBackupVaultsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListBackupVaultsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListBackupVaultsWithContext(ctx context.Context, input *backup.ListBackupVaultsInput, opts ...request.Option) (*backup.ListBackupVaultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListBackupVaultsOutput), nil
	}
	out, err := c.BackupAPI.ListBackupVaultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListCopyJobs(input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListCopyJobsOutput), nil
	}
	out, err := c.BackupAPI.ListCopyJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListCopyJobsPages(input *backup.ListCopyJobsInput, fn func(*backup.ListCopyJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListCopyJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListCopyJobsOutput{}
	fnCacher := func(out *backup.ListCopyJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListCopyJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListCopyJobsPagesWithContext(ctx context.Context, input *backup.ListCopyJobsInput, fn func(*backup.ListCopyJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListCopyJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListCopyJobsOutput{}
	fnCacher := func(out *backup.ListCopyJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListCopyJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListCopyJobsWithContext(ctx context.Context, input *backup.ListCopyJobsInput, opts ...request.Option) (*backup.ListCopyJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListCopyJobsOutput), nil
	}
	out, err := c.BackupAPI.ListCopyJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListFrameworks(input *backup.ListFrameworksInput) (*backup.ListFrameworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListFrameworksOutput), nil
	}
	out, err := c.BackupAPI.ListFrameworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListFrameworksPages(input *backup.ListFrameworksInput, fn func(*backup.ListFrameworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListFrameworksOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListFrameworksOutput{}
	fnCacher := func(out *backup.ListFrameworksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListFrameworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListFrameworksPagesWithContext(ctx context.Context, input *backup.ListFrameworksInput, fn func(*backup.ListFrameworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListFrameworksOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListFrameworksOutput{}
	fnCacher := func(out *backup.ListFrameworksOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListFrameworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListFrameworksWithContext(ctx context.Context, input *backup.ListFrameworksInput, opts ...request.Option) (*backup.ListFrameworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListFrameworksOutput), nil
	}
	out, err := c.BackupAPI.ListFrameworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListLegalHolds(input *backup.ListLegalHoldsInput) (*backup.ListLegalHoldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListLegalHoldsOutput), nil
	}
	out, err := c.BackupAPI.ListLegalHolds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListLegalHoldsPages(input *backup.ListLegalHoldsInput, fn func(*backup.ListLegalHoldsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListLegalHoldsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListLegalHoldsOutput{}
	fnCacher := func(out *backup.ListLegalHoldsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListLegalHoldsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListLegalHoldsPagesWithContext(ctx context.Context, input *backup.ListLegalHoldsInput, fn func(*backup.ListLegalHoldsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListLegalHoldsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListLegalHoldsOutput{}
	fnCacher := func(out *backup.ListLegalHoldsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListLegalHoldsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListLegalHoldsWithContext(ctx context.Context, input *backup.ListLegalHoldsInput, opts ...request.Option) (*backup.ListLegalHoldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListLegalHoldsOutput), nil
	}
	out, err := c.BackupAPI.ListLegalHoldsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListProtectedResources(input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListProtectedResourcesOutput), nil
	}
	out, err := c.BackupAPI.ListProtectedResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListProtectedResourcesPages(input *backup.ListProtectedResourcesInput, fn func(*backup.ListProtectedResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListProtectedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListProtectedResourcesOutput{}
	fnCacher := func(out *backup.ListProtectedResourcesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListProtectedResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListProtectedResourcesPagesWithContext(ctx context.Context, input *backup.ListProtectedResourcesInput, fn func(*backup.ListProtectedResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListProtectedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListProtectedResourcesOutput{}
	fnCacher := func(out *backup.ListProtectedResourcesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListProtectedResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListProtectedResourcesWithContext(ctx context.Context, input *backup.ListProtectedResourcesInput, opts ...request.Option) (*backup.ListProtectedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListProtectedResourcesOutput), nil
	}
	out, err := c.BackupAPI.ListProtectedResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRecoveryPointsByBackupVault(input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRecoveryPointsByBackupVaultOutput), nil
	}
	out, err := c.BackupAPI.ListRecoveryPointsByBackupVault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRecoveryPointsByBackupVaultPages(input *backup.ListRecoveryPointsByBackupVaultInput, fn func(*backup.ListRecoveryPointsByBackupVaultOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRecoveryPointsByBackupVaultOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRecoveryPointsByBackupVaultOutput{}
	fnCacher := func(out *backup.ListRecoveryPointsByBackupVaultOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRecoveryPointsByBackupVaultPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRecoveryPointsByBackupVaultPagesWithContext(ctx context.Context, input *backup.ListRecoveryPointsByBackupVaultInput, fn func(*backup.ListRecoveryPointsByBackupVaultOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRecoveryPointsByBackupVaultOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRecoveryPointsByBackupVaultOutput{}
	fnCacher := func(out *backup.ListRecoveryPointsByBackupVaultOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRecoveryPointsByBackupVaultPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRecoveryPointsByBackupVaultWithContext(ctx context.Context, input *backup.ListRecoveryPointsByBackupVaultInput, opts ...request.Option) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRecoveryPointsByBackupVaultOutput), nil
	}
	out, err := c.BackupAPI.ListRecoveryPointsByBackupVaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRecoveryPointsByLegalHold(input *backup.ListRecoveryPointsByLegalHoldInput) (*backup.ListRecoveryPointsByLegalHoldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRecoveryPointsByLegalHoldOutput), nil
	}
	out, err := c.BackupAPI.ListRecoveryPointsByLegalHold(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRecoveryPointsByLegalHoldPages(input *backup.ListRecoveryPointsByLegalHoldInput, fn func(*backup.ListRecoveryPointsByLegalHoldOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRecoveryPointsByLegalHoldOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRecoveryPointsByLegalHoldOutput{}
	fnCacher := func(out *backup.ListRecoveryPointsByLegalHoldOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRecoveryPointsByLegalHoldPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRecoveryPointsByLegalHoldPagesWithContext(ctx context.Context, input *backup.ListRecoveryPointsByLegalHoldInput, fn func(*backup.ListRecoveryPointsByLegalHoldOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRecoveryPointsByLegalHoldOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRecoveryPointsByLegalHoldOutput{}
	fnCacher := func(out *backup.ListRecoveryPointsByLegalHoldOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRecoveryPointsByLegalHoldPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRecoveryPointsByLegalHoldWithContext(ctx context.Context, input *backup.ListRecoveryPointsByLegalHoldInput, opts ...request.Option) (*backup.ListRecoveryPointsByLegalHoldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRecoveryPointsByLegalHoldOutput), nil
	}
	out, err := c.BackupAPI.ListRecoveryPointsByLegalHoldWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRecoveryPointsByResource(input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRecoveryPointsByResourceOutput), nil
	}
	out, err := c.BackupAPI.ListRecoveryPointsByResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRecoveryPointsByResourcePages(input *backup.ListRecoveryPointsByResourceInput, fn func(*backup.ListRecoveryPointsByResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRecoveryPointsByResourceOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRecoveryPointsByResourceOutput{}
	fnCacher := func(out *backup.ListRecoveryPointsByResourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRecoveryPointsByResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRecoveryPointsByResourcePagesWithContext(ctx context.Context, input *backup.ListRecoveryPointsByResourceInput, fn func(*backup.ListRecoveryPointsByResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRecoveryPointsByResourceOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRecoveryPointsByResourceOutput{}
	fnCacher := func(out *backup.ListRecoveryPointsByResourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRecoveryPointsByResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRecoveryPointsByResourceWithContext(ctx context.Context, input *backup.ListRecoveryPointsByResourceInput, opts ...request.Option) (*backup.ListRecoveryPointsByResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRecoveryPointsByResourceOutput), nil
	}
	out, err := c.BackupAPI.ListRecoveryPointsByResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListReportJobs(input *backup.ListReportJobsInput) (*backup.ListReportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListReportJobsOutput), nil
	}
	out, err := c.BackupAPI.ListReportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListReportJobsPages(input *backup.ListReportJobsInput, fn func(*backup.ListReportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListReportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListReportJobsOutput{}
	fnCacher := func(out *backup.ListReportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListReportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListReportJobsPagesWithContext(ctx context.Context, input *backup.ListReportJobsInput, fn func(*backup.ListReportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListReportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListReportJobsOutput{}
	fnCacher := func(out *backup.ListReportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListReportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListReportJobsWithContext(ctx context.Context, input *backup.ListReportJobsInput, opts ...request.Option) (*backup.ListReportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListReportJobsOutput), nil
	}
	out, err := c.BackupAPI.ListReportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListReportPlans(input *backup.ListReportPlansInput) (*backup.ListReportPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListReportPlansOutput), nil
	}
	out, err := c.BackupAPI.ListReportPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListReportPlansPages(input *backup.ListReportPlansInput, fn func(*backup.ListReportPlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListReportPlansOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListReportPlansOutput{}
	fnCacher := func(out *backup.ListReportPlansOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListReportPlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListReportPlansPagesWithContext(ctx context.Context, input *backup.ListReportPlansInput, fn func(*backup.ListReportPlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListReportPlansOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListReportPlansOutput{}
	fnCacher := func(out *backup.ListReportPlansOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListReportPlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListReportPlansWithContext(ctx context.Context, input *backup.ListReportPlansInput, opts ...request.Option) (*backup.ListReportPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListReportPlansOutput), nil
	}
	out, err := c.BackupAPI.ListReportPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRestoreJobs(input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRestoreJobsOutput), nil
	}
	out, err := c.BackupAPI.ListRestoreJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListRestoreJobsPages(input *backup.ListRestoreJobsInput, fn func(*backup.ListRestoreJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRestoreJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRestoreJobsOutput{}
	fnCacher := func(out *backup.ListRestoreJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRestoreJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRestoreJobsPagesWithContext(ctx context.Context, input *backup.ListRestoreJobsInput, fn func(*backup.ListRestoreJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListRestoreJobsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListRestoreJobsOutput{}
	fnCacher := func(out *backup.ListRestoreJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListRestoreJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListRestoreJobsWithContext(ctx context.Context, input *backup.ListRestoreJobsInput, opts ...request.Option) (*backup.ListRestoreJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListRestoreJobsOutput), nil
	}
	out, err := c.BackupAPI.ListRestoreJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListTags(input *backup.ListTagsInput) (*backup.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListTagsOutput), nil
	}
	out, err := c.BackupAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Backup) ListTagsPages(input *backup.ListTagsInput, fn func(*backup.ListTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListTagsOutput{}
	fnCacher := func(out *backup.ListTagsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListTagsPagesWithContext(ctx context.Context, input *backup.ListTagsInput, fn func(*backup.ListTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backup.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &backup.ListTagsOutput{}
	fnCacher := func(out *backup.ListTagsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.BackupAPI.ListTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Backup) ListTagsWithContext(ctx context.Context, input *backup.ListTagsInput, opts ...request.Option) (*backup.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backup.ListTagsOutput), nil
	}
	out, err := c.BackupAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

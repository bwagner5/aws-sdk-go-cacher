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
package customerprofilescacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/customerprofiles"
	"github.com/aws/aws-sdk-go/service/customerprofiles/customerprofilesiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CustomerProfiles struct {
	customerprofilesiface.CustomerProfilesAPI
	cache *cache.Cache
}

func New(customerprofilesapi customerprofilesiface.CustomerProfilesAPI) *CustomerProfiles {
	return &CustomerProfiles{
		CustomerProfilesAPI: customerprofilesapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CustomerProfiles) GetAutoMergingPreview(input *customerprofiles.GetAutoMergingPreviewInput) (*customerprofiles.GetAutoMergingPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetAutoMergingPreviewOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetAutoMergingPreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetAutoMergingPreviewWithContext(ctx context.Context, input *customerprofiles.GetAutoMergingPreviewInput, opts ...request.Option) (*customerprofiles.GetAutoMergingPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetAutoMergingPreviewOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetAutoMergingPreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetDomain(input *customerprofiles.GetDomainInput) (*customerprofiles.GetDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetDomainOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetDomainWithContext(ctx context.Context, input *customerprofiles.GetDomainInput, opts ...request.Option) (*customerprofiles.GetDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetDomainOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetIdentityResolutionJob(input *customerprofiles.GetIdentityResolutionJobInput) (*customerprofiles.GetIdentityResolutionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetIdentityResolutionJobOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetIdentityResolutionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetIdentityResolutionJobWithContext(ctx context.Context, input *customerprofiles.GetIdentityResolutionJobInput, opts ...request.Option) (*customerprofiles.GetIdentityResolutionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetIdentityResolutionJobOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetIdentityResolutionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetIntegration(input *customerprofiles.GetIntegrationInput) (*customerprofiles.GetIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetIntegrationOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetIntegration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetIntegrationWithContext(ctx context.Context, input *customerprofiles.GetIntegrationInput, opts ...request.Option) (*customerprofiles.GetIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetIntegrationOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetIntegrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetMatches(input *customerprofiles.GetMatchesInput) (*customerprofiles.GetMatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetMatchesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetMatches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetMatchesWithContext(ctx context.Context, input *customerprofiles.GetMatchesInput, opts ...request.Option) (*customerprofiles.GetMatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetMatchesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetMatchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetProfileObjectType(input *customerprofiles.GetProfileObjectTypeInput) (*customerprofiles.GetProfileObjectTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetProfileObjectTypeOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetProfileObjectType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetProfileObjectTypeTemplate(input *customerprofiles.GetProfileObjectTypeTemplateInput) (*customerprofiles.GetProfileObjectTypeTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetProfileObjectTypeTemplateOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetProfileObjectTypeTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetProfileObjectTypeTemplateWithContext(ctx context.Context, input *customerprofiles.GetProfileObjectTypeTemplateInput, opts ...request.Option) (*customerprofiles.GetProfileObjectTypeTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetProfileObjectTypeTemplateOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetProfileObjectTypeTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetProfileObjectTypeWithContext(ctx context.Context, input *customerprofiles.GetProfileObjectTypeInput, opts ...request.Option) (*customerprofiles.GetProfileObjectTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetProfileObjectTypeOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetProfileObjectTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetWorkflow(input *customerprofiles.GetWorkflowInput) (*customerprofiles.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetWorkflowOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetWorkflow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetWorkflowSteps(input *customerprofiles.GetWorkflowStepsInput) (*customerprofiles.GetWorkflowStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetWorkflowStepsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetWorkflowSteps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetWorkflowStepsWithContext(ctx context.Context, input *customerprofiles.GetWorkflowStepsInput, opts ...request.Option) (*customerprofiles.GetWorkflowStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetWorkflowStepsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetWorkflowStepsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) GetWorkflowWithContext(ctx context.Context, input *customerprofiles.GetWorkflowInput, opts ...request.Option) (*customerprofiles.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.GetWorkflowOutput), nil
	}
	out, err := c.CustomerProfilesAPI.GetWorkflowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListAccountIntegrations(input *customerprofiles.ListAccountIntegrationsInput) (*customerprofiles.ListAccountIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListAccountIntegrationsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListAccountIntegrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListAccountIntegrationsWithContext(ctx context.Context, input *customerprofiles.ListAccountIntegrationsInput, opts ...request.Option) (*customerprofiles.ListAccountIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListAccountIntegrationsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListAccountIntegrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListDomains(input *customerprofiles.ListDomainsInput) (*customerprofiles.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListDomainsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListDomainsWithContext(ctx context.Context, input *customerprofiles.ListDomainsInput, opts ...request.Option) (*customerprofiles.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListDomainsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListIdentityResolutionJobs(input *customerprofiles.ListIdentityResolutionJobsInput) (*customerprofiles.ListIdentityResolutionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListIdentityResolutionJobsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListIdentityResolutionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListIdentityResolutionJobsWithContext(ctx context.Context, input *customerprofiles.ListIdentityResolutionJobsInput, opts ...request.Option) (*customerprofiles.ListIdentityResolutionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListIdentityResolutionJobsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListIdentityResolutionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListIntegrations(input *customerprofiles.ListIntegrationsInput) (*customerprofiles.ListIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListIntegrationsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListIntegrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListIntegrationsWithContext(ctx context.Context, input *customerprofiles.ListIntegrationsInput, opts ...request.Option) (*customerprofiles.ListIntegrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListIntegrationsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListIntegrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListProfileObjectTypeTemplates(input *customerprofiles.ListProfileObjectTypeTemplatesInput) (*customerprofiles.ListProfileObjectTypeTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListProfileObjectTypeTemplatesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListProfileObjectTypeTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListProfileObjectTypeTemplatesWithContext(ctx context.Context, input *customerprofiles.ListProfileObjectTypeTemplatesInput, opts ...request.Option) (*customerprofiles.ListProfileObjectTypeTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListProfileObjectTypeTemplatesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListProfileObjectTypeTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListProfileObjectTypes(input *customerprofiles.ListProfileObjectTypesInput) (*customerprofiles.ListProfileObjectTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListProfileObjectTypesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListProfileObjectTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListProfileObjectTypesWithContext(ctx context.Context, input *customerprofiles.ListProfileObjectTypesInput, opts ...request.Option) (*customerprofiles.ListProfileObjectTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListProfileObjectTypesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListProfileObjectTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListProfileObjects(input *customerprofiles.ListProfileObjectsInput) (*customerprofiles.ListProfileObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListProfileObjectsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListProfileObjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListProfileObjectsWithContext(ctx context.Context, input *customerprofiles.ListProfileObjectsInput, opts ...request.Option) (*customerprofiles.ListProfileObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListProfileObjectsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListProfileObjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListTagsForResource(input *customerprofiles.ListTagsForResourceInput) (*customerprofiles.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListTagsForResourceOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListTagsForResourceWithContext(ctx context.Context, input *customerprofiles.ListTagsForResourceInput, opts ...request.Option) (*customerprofiles.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListTagsForResourceOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListWorkflows(input *customerprofiles.ListWorkflowsInput) (*customerprofiles.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListWorkflowsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListWorkflows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) ListWorkflowsWithContext(ctx context.Context, input *customerprofiles.ListWorkflowsInput, opts ...request.Option) (*customerprofiles.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.ListWorkflowsOutput), nil
	}
	out, err := c.CustomerProfilesAPI.ListWorkflowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) SearchProfiles(input *customerprofiles.SearchProfilesInput) (*customerprofiles.SearchProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.SearchProfilesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.SearchProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CustomerProfiles) SearchProfilesWithContext(ctx context.Context, input *customerprofiles.SearchProfilesInput, opts ...request.Option) (*customerprofiles.SearchProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*customerprofiles.SearchProfilesOutput), nil
	}
	out, err := c.CustomerProfilesAPI.SearchProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

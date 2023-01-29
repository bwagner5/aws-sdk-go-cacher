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
package amplifycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Amplify struct {
	amplifyiface.AmplifyAPI
	cache *cache.Cache
}

func New(amplifyapi amplifyiface.AmplifyAPI) *Amplify {
	return &Amplify{
		AmplifyAPI: amplifyapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Amplify) GetApp(input *amplify.GetAppInput) (*amplify.GetAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetAppOutput), nil
	}
	out, err := c.AmplifyAPI.GetApp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetAppWithContext(ctx context.Context, input *amplify.GetAppInput, opts ...request.Option) (*amplify.GetAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetAppOutput), nil
	}
	out, err := c.AmplifyAPI.GetAppWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetArtifactUrl(input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetArtifactUrlOutput), nil
	}
	out, err := c.AmplifyAPI.GetArtifactUrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetArtifactUrlWithContext(ctx context.Context, input *amplify.GetArtifactUrlInput, opts ...request.Option) (*amplify.GetArtifactUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetArtifactUrlOutput), nil
	}
	out, err := c.AmplifyAPI.GetArtifactUrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetBackendEnvironment(input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetBackendEnvironmentOutput), nil
	}
	out, err := c.AmplifyAPI.GetBackendEnvironment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetBackendEnvironmentWithContext(ctx context.Context, input *amplify.GetBackendEnvironmentInput, opts ...request.Option) (*amplify.GetBackendEnvironmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetBackendEnvironmentOutput), nil
	}
	out, err := c.AmplifyAPI.GetBackendEnvironmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetBranch(input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetBranchOutput), nil
	}
	out, err := c.AmplifyAPI.GetBranch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetBranchWithContext(ctx context.Context, input *amplify.GetBranchInput, opts ...request.Option) (*amplify.GetBranchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetBranchOutput), nil
	}
	out, err := c.AmplifyAPI.GetBranchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetDomainAssociation(input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetDomainAssociationOutput), nil
	}
	out, err := c.AmplifyAPI.GetDomainAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetDomainAssociationWithContext(ctx context.Context, input *amplify.GetDomainAssociationInput, opts ...request.Option) (*amplify.GetDomainAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetDomainAssociationOutput), nil
	}
	out, err := c.AmplifyAPI.GetDomainAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetJob(input *amplify.GetJobInput) (*amplify.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetJobOutput), nil
	}
	out, err := c.AmplifyAPI.GetJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetJobWithContext(ctx context.Context, input *amplify.GetJobInput, opts ...request.Option) (*amplify.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetJobOutput), nil
	}
	out, err := c.AmplifyAPI.GetJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetWebhook(input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetWebhookOutput), nil
	}
	out, err := c.AmplifyAPI.GetWebhook(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) GetWebhookWithContext(ctx context.Context, input *amplify.GetWebhookInput, opts ...request.Option) (*amplify.GetWebhookOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.GetWebhookOutput), nil
	}
	out, err := c.AmplifyAPI.GetWebhookWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListApps(input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListAppsOutput), nil
	}
	out, err := c.AmplifyAPI.ListApps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListAppsWithContext(ctx context.Context, input *amplify.ListAppsInput, opts ...request.Option) (*amplify.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListAppsOutput), nil
	}
	out, err := c.AmplifyAPI.ListAppsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListArtifacts(input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListArtifactsOutput), nil
	}
	out, err := c.AmplifyAPI.ListArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListArtifactsWithContext(ctx context.Context, input *amplify.ListArtifactsInput, opts ...request.Option) (*amplify.ListArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListArtifactsOutput), nil
	}
	out, err := c.AmplifyAPI.ListArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListBackendEnvironments(input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListBackendEnvironmentsOutput), nil
	}
	out, err := c.AmplifyAPI.ListBackendEnvironments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListBackendEnvironmentsWithContext(ctx context.Context, input *amplify.ListBackendEnvironmentsInput, opts ...request.Option) (*amplify.ListBackendEnvironmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListBackendEnvironmentsOutput), nil
	}
	out, err := c.AmplifyAPI.ListBackendEnvironmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListBranches(input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListBranchesOutput), nil
	}
	out, err := c.AmplifyAPI.ListBranches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListBranchesWithContext(ctx context.Context, input *amplify.ListBranchesInput, opts ...request.Option) (*amplify.ListBranchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListBranchesOutput), nil
	}
	out, err := c.AmplifyAPI.ListBranchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListDomainAssociations(input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListDomainAssociationsOutput), nil
	}
	out, err := c.AmplifyAPI.ListDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListDomainAssociationsWithContext(ctx context.Context, input *amplify.ListDomainAssociationsInput, opts ...request.Option) (*amplify.ListDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListDomainAssociationsOutput), nil
	}
	out, err := c.AmplifyAPI.ListDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListJobs(input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListJobsOutput), nil
	}
	out, err := c.AmplifyAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListJobsWithContext(ctx context.Context, input *amplify.ListJobsInput, opts ...request.Option) (*amplify.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListJobsOutput), nil
	}
	out, err := c.AmplifyAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListTagsForResource(input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListTagsForResourceOutput), nil
	}
	out, err := c.AmplifyAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListTagsForResourceWithContext(ctx context.Context, input *amplify.ListTagsForResourceInput, opts ...request.Option) (*amplify.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListTagsForResourceOutput), nil
	}
	out, err := c.AmplifyAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListWebhooks(input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListWebhooksOutput), nil
	}
	out, err := c.AmplifyAPI.ListWebhooks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Amplify) ListWebhooksWithContext(ctx context.Context, input *amplify.ListWebhooksInput, opts ...request.Option) (*amplify.ListWebhooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*amplify.ListWebhooksOutput), nil
	}
	out, err := c.AmplifyAPI.ListWebhooksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

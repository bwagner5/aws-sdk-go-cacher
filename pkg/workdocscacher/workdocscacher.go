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
package workdocscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/workdocs/workdocsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WorkDocs struct {
	workdocsiface.WorkDocsAPI
	cache *cache.Cache
}

func New(workdocsapi workdocsiface.WorkDocsAPI) *WorkDocs {
	return &WorkDocs{
		WorkDocsAPI: workdocsapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WorkDocs) DescribeActivities(input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeActivitiesOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeActivities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeActivitiesWithContext(ctx context.Context, input *workdocs.DescribeActivitiesInput, opts ...request.Option) (*workdocs.DescribeActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeActivitiesOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeActivitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeComments(input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeCommentsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeComments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeCommentsWithContext(ctx context.Context, input *workdocs.DescribeCommentsInput, opts ...request.Option) (*workdocs.DescribeCommentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeCommentsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeCommentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeDocumentVersions(input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeDocumentVersionsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeDocumentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeDocumentVersionsPages(input *workdocs.DescribeDocumentVersionsInput, fn func(*workdocs.DescribeDocumentVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workdocs.DescribeDocumentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &workdocs.DescribeDocumentVersionsOutput{}
	fnCacher := func(out *workdocs.DescribeDocumentVersionsOutput, more bool) bool {
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
	if err := c.WorkDocsAPI.DescribeDocumentVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkDocs) DescribeDocumentVersionsPagesWithContext(ctx context.Context, input *workdocs.DescribeDocumentVersionsInput, fn func(*workdocs.DescribeDocumentVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workdocs.DescribeDocumentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &workdocs.DescribeDocumentVersionsOutput{}
	fnCacher := func(out *workdocs.DescribeDocumentVersionsOutput, more bool) bool {
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
	if err := c.WorkDocsAPI.DescribeDocumentVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkDocs) DescribeDocumentVersionsWithContext(ctx context.Context, input *workdocs.DescribeDocumentVersionsInput, opts ...request.Option) (*workdocs.DescribeDocumentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeDocumentVersionsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeDocumentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeFolderContents(input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeFolderContentsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeFolderContents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeFolderContentsPages(input *workdocs.DescribeFolderContentsInput, fn func(*workdocs.DescribeFolderContentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workdocs.DescribeFolderContentsOutput), false)
		return nil
	}
	cachable := true
	output := &workdocs.DescribeFolderContentsOutput{}
	fnCacher := func(out *workdocs.DescribeFolderContentsOutput, more bool) bool {
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
	if err := c.WorkDocsAPI.DescribeFolderContentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkDocs) DescribeFolderContentsPagesWithContext(ctx context.Context, input *workdocs.DescribeFolderContentsInput, fn func(*workdocs.DescribeFolderContentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workdocs.DescribeFolderContentsOutput), false)
		return nil
	}
	cachable := true
	output := &workdocs.DescribeFolderContentsOutput{}
	fnCacher := func(out *workdocs.DescribeFolderContentsOutput, more bool) bool {
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
	if err := c.WorkDocsAPI.DescribeFolderContentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkDocs) DescribeFolderContentsWithContext(ctx context.Context, input *workdocs.DescribeFolderContentsInput, opts ...request.Option) (*workdocs.DescribeFolderContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeFolderContentsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeFolderContentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeGroups(input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeGroupsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeGroupsWithContext(ctx context.Context, input *workdocs.DescribeGroupsInput, opts ...request.Option) (*workdocs.DescribeGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeGroupsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeNotificationSubscriptions(input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeNotificationSubscriptionsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeNotificationSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeNotificationSubscriptionsWithContext(ctx context.Context, input *workdocs.DescribeNotificationSubscriptionsInput, opts ...request.Option) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeNotificationSubscriptionsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeNotificationSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeResourcePermissions(input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeResourcePermissionsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeResourcePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeResourcePermissionsWithContext(ctx context.Context, input *workdocs.DescribeResourcePermissionsInput, opts ...request.Option) (*workdocs.DescribeResourcePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeResourcePermissionsOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeResourcePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeRootFolders(input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeRootFoldersOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeRootFolders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeRootFoldersWithContext(ctx context.Context, input *workdocs.DescribeRootFoldersInput, opts ...request.Option) (*workdocs.DescribeRootFoldersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeRootFoldersOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeRootFoldersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeUsers(input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeUsersOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) DescribeUsersPages(input *workdocs.DescribeUsersInput, fn func(*workdocs.DescribeUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workdocs.DescribeUsersOutput), false)
		return nil
	}
	cachable := true
	output := &workdocs.DescribeUsersOutput{}
	fnCacher := func(out *workdocs.DescribeUsersOutput, more bool) bool {
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
	if err := c.WorkDocsAPI.DescribeUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkDocs) DescribeUsersPagesWithContext(ctx context.Context, input *workdocs.DescribeUsersInput, fn func(*workdocs.DescribeUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workdocs.DescribeUsersOutput), false)
		return nil
	}
	cachable := true
	output := &workdocs.DescribeUsersOutput{}
	fnCacher := func(out *workdocs.DescribeUsersOutput, more bool) bool {
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
	if err := c.WorkDocsAPI.DescribeUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkDocs) DescribeUsersWithContext(ctx context.Context, input *workdocs.DescribeUsersInput, opts ...request.Option) (*workdocs.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.DescribeUsersOutput), nil
	}
	out, err := c.WorkDocsAPI.DescribeUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetCurrentUser(input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetCurrentUserOutput), nil
	}
	out, err := c.WorkDocsAPI.GetCurrentUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetCurrentUserWithContext(ctx context.Context, input *workdocs.GetCurrentUserInput, opts ...request.Option) (*workdocs.GetCurrentUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetCurrentUserOutput), nil
	}
	out, err := c.WorkDocsAPI.GetCurrentUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetDocument(input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetDocumentOutput), nil
	}
	out, err := c.WorkDocsAPI.GetDocument(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetDocumentPath(input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetDocumentPathOutput), nil
	}
	out, err := c.WorkDocsAPI.GetDocumentPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetDocumentPathWithContext(ctx context.Context, input *workdocs.GetDocumentPathInput, opts ...request.Option) (*workdocs.GetDocumentPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetDocumentPathOutput), nil
	}
	out, err := c.WorkDocsAPI.GetDocumentPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetDocumentVersion(input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetDocumentVersionOutput), nil
	}
	out, err := c.WorkDocsAPI.GetDocumentVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetDocumentVersionWithContext(ctx context.Context, input *workdocs.GetDocumentVersionInput, opts ...request.Option) (*workdocs.GetDocumentVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetDocumentVersionOutput), nil
	}
	out, err := c.WorkDocsAPI.GetDocumentVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetDocumentWithContext(ctx context.Context, input *workdocs.GetDocumentInput, opts ...request.Option) (*workdocs.GetDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetDocumentOutput), nil
	}
	out, err := c.WorkDocsAPI.GetDocumentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetFolder(input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetFolderOutput), nil
	}
	out, err := c.WorkDocsAPI.GetFolder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetFolderPath(input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetFolderPathOutput), nil
	}
	out, err := c.WorkDocsAPI.GetFolderPath(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetFolderPathWithContext(ctx context.Context, input *workdocs.GetFolderPathInput, opts ...request.Option) (*workdocs.GetFolderPathOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetFolderPathOutput), nil
	}
	out, err := c.WorkDocsAPI.GetFolderPathWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetFolderWithContext(ctx context.Context, input *workdocs.GetFolderInput, opts ...request.Option) (*workdocs.GetFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetFolderOutput), nil
	}
	out, err := c.WorkDocsAPI.GetFolderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetResources(input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetResourcesOutput), nil
	}
	out, err := c.WorkDocsAPI.GetResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkDocs) GetResourcesWithContext(ctx context.Context, input *workdocs.GetResourcesInput, opts ...request.Option) (*workdocs.GetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workdocs.GetResourcesOutput), nil
	}
	out, err := c.WorkDocsAPI.GetResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

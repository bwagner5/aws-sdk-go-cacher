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
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codecommit/codecommitiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeCommit struct {
	codecommitiface.CodeCommitAPI
	cache *cache.Cache
}

func NewCodeCommit(codecommitapi codecommitiface.CodeCommitAPI) *CodeCommit {
	return &CodeCommit{
		CodeCommitAPI: codecommitapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeCommit) BatchAssociateApprovalRuleTemplateWithRepositories(input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchAssociateApprovalRuleTemplateWithRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchAssociateApprovalRuleTemplateWithRepositoriesWithContext(ctx context.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput, opts ...request.Option) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchAssociateApprovalRuleTemplateWithRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchDescribeMergeConflicts(input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchDescribeMergeConflictsOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchDescribeMergeConflicts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchDescribeMergeConflictsWithContext(ctx context.Context, input *codecommit.BatchDescribeMergeConflictsInput, opts ...request.Option) (*codecommit.BatchDescribeMergeConflictsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchDescribeMergeConflictsOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchDescribeMergeConflictsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchDisassociateApprovalRuleTemplateFromRepositories(input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchDisassociateApprovalRuleTemplateFromRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchDisassociateApprovalRuleTemplateFromRepositoriesWithContext(ctx context.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput, opts ...request.Option) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchDisassociateApprovalRuleTemplateFromRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchGetCommits(input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchGetCommitsOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchGetCommits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchGetCommitsWithContext(ctx context.Context, input *codecommit.BatchGetCommitsInput, opts ...request.Option) (*codecommit.BatchGetCommitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchGetCommitsOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchGetCommitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchGetRepositories(input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchGetRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchGetRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) BatchGetRepositoriesWithContext(ctx context.Context, input *codecommit.BatchGetRepositoriesInput, opts ...request.Option) (*codecommit.BatchGetRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.BatchGetRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.BatchGetRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) DescribeMergeConflicts(input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.DescribeMergeConflictsOutput), nil
	}
	out, err := c.CodeCommitAPI.DescribeMergeConflicts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) DescribeMergeConflictsPages(input *codecommit.DescribeMergeConflictsInput, fn func(*codecommit.DescribeMergeConflictsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.DescribeMergeConflictsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.DescribeMergeConflictsOutput{}
	fnCacher := func(out *codecommit.DescribeMergeConflictsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.DescribeMergeConflictsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) DescribeMergeConflictsPagesWithContext(ctx context.Context, input *codecommit.DescribeMergeConflictsInput, fn func(*codecommit.DescribeMergeConflictsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.DescribeMergeConflictsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.DescribeMergeConflictsOutput{}
	fnCacher := func(out *codecommit.DescribeMergeConflictsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.DescribeMergeConflictsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) DescribeMergeConflictsWithContext(ctx context.Context, input *codecommit.DescribeMergeConflictsInput, opts ...request.Option) (*codecommit.DescribeMergeConflictsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.DescribeMergeConflictsOutput), nil
	}
	out, err := c.CodeCommitAPI.DescribeMergeConflictsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) DescribePullRequestEvents(input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.DescribePullRequestEventsOutput), nil
	}
	out, err := c.CodeCommitAPI.DescribePullRequestEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) DescribePullRequestEventsPages(input *codecommit.DescribePullRequestEventsInput, fn func(*codecommit.DescribePullRequestEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.DescribePullRequestEventsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.DescribePullRequestEventsOutput{}
	fnCacher := func(out *codecommit.DescribePullRequestEventsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.DescribePullRequestEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) DescribePullRequestEventsPagesWithContext(ctx context.Context, input *codecommit.DescribePullRequestEventsInput, fn func(*codecommit.DescribePullRequestEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.DescribePullRequestEventsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.DescribePullRequestEventsOutput{}
	fnCacher := func(out *codecommit.DescribePullRequestEventsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.DescribePullRequestEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) DescribePullRequestEventsWithContext(ctx context.Context, input *codecommit.DescribePullRequestEventsInput, opts ...request.Option) (*codecommit.DescribePullRequestEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.DescribePullRequestEventsOutput), nil
	}
	out, err := c.CodeCommitAPI.DescribePullRequestEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetApprovalRuleTemplate(input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetApprovalRuleTemplateOutput), nil
	}
	out, err := c.CodeCommitAPI.GetApprovalRuleTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetApprovalRuleTemplateWithContext(ctx context.Context, input *codecommit.GetApprovalRuleTemplateInput, opts ...request.Option) (*codecommit.GetApprovalRuleTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetApprovalRuleTemplateOutput), nil
	}
	out, err := c.CodeCommitAPI.GetApprovalRuleTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetBlob(input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetBlobOutput), nil
	}
	out, err := c.CodeCommitAPI.GetBlob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetBlobWithContext(ctx context.Context, input *codecommit.GetBlobInput, opts ...request.Option) (*codecommit.GetBlobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetBlobOutput), nil
	}
	out, err := c.CodeCommitAPI.GetBlobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetBranch(input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetBranchOutput), nil
	}
	out, err := c.CodeCommitAPI.GetBranch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetBranchWithContext(ctx context.Context, input *codecommit.GetBranchInput, opts ...request.Option) (*codecommit.GetBranchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetBranchOutput), nil
	}
	out, err := c.CodeCommitAPI.GetBranchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetComment(input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentOutput), nil
	}
	out, err := c.CodeCommitAPI.GetComment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommentReactions(input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentReactionsOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommentReactions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommentReactionsPages(input *codecommit.GetCommentReactionsInput, fn func(*codecommit.GetCommentReactionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetCommentReactionsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetCommentReactionsOutput{}
	fnCacher := func(out *codecommit.GetCommentReactionsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetCommentReactionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetCommentReactionsPagesWithContext(ctx context.Context, input *codecommit.GetCommentReactionsInput, fn func(*codecommit.GetCommentReactionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetCommentReactionsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetCommentReactionsOutput{}
	fnCacher := func(out *codecommit.GetCommentReactionsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetCommentReactionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetCommentReactionsWithContext(ctx context.Context, input *codecommit.GetCommentReactionsInput, opts ...request.Option) (*codecommit.GetCommentReactionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentReactionsOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommentReactionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommentWithContext(ctx context.Context, input *codecommit.GetCommentInput, opts ...request.Option) (*codecommit.GetCommentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommentsForComparedCommit(input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentsForComparedCommitOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommentsForComparedCommit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommentsForComparedCommitPages(input *codecommit.GetCommentsForComparedCommitInput, fn func(*codecommit.GetCommentsForComparedCommitOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetCommentsForComparedCommitOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetCommentsForComparedCommitOutput{}
	fnCacher := func(out *codecommit.GetCommentsForComparedCommitOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetCommentsForComparedCommitPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetCommentsForComparedCommitPagesWithContext(ctx context.Context, input *codecommit.GetCommentsForComparedCommitInput, fn func(*codecommit.GetCommentsForComparedCommitOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetCommentsForComparedCommitOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetCommentsForComparedCommitOutput{}
	fnCacher := func(out *codecommit.GetCommentsForComparedCommitOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetCommentsForComparedCommitPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetCommentsForComparedCommitWithContext(ctx context.Context, input *codecommit.GetCommentsForComparedCommitInput, opts ...request.Option) (*codecommit.GetCommentsForComparedCommitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentsForComparedCommitOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommentsForComparedCommitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommentsForPullRequestPages(input *codecommit.GetCommentsForPullRequestInput, fn func(*codecommit.GetCommentsForPullRequestOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetCommentsForPullRequestOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetCommentsForPullRequestOutput{}
	fnCacher := func(out *codecommit.GetCommentsForPullRequestOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetCommentsForPullRequestPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetCommentsForPullRequestPagesWithContext(ctx context.Context, input *codecommit.GetCommentsForPullRequestInput, fn func(*codecommit.GetCommentsForPullRequestOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetCommentsForPullRequestOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetCommentsForPullRequestOutput{}
	fnCacher := func(out *codecommit.GetCommentsForPullRequestOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetCommentsForPullRequestPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetCommentsForPullRequestWithContext(ctx context.Context, input *codecommit.GetCommentsForPullRequestInput, opts ...request.Option) (*codecommit.GetCommentsForPullRequestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommentsForPullRequestOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommentsForPullRequestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommit(input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommitOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetCommitWithContext(ctx context.Context, input *codecommit.GetCommitInput, opts ...request.Option) (*codecommit.GetCommitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetCommitOutput), nil
	}
	out, err := c.CodeCommitAPI.GetCommitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetDifferences(input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetDifferencesOutput), nil
	}
	out, err := c.CodeCommitAPI.GetDifferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetDifferencesPages(input *codecommit.GetDifferencesInput, fn func(*codecommit.GetDifferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetDifferencesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetDifferencesOutput{}
	fnCacher := func(out *codecommit.GetDifferencesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetDifferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetDifferencesPagesWithContext(ctx context.Context, input *codecommit.GetDifferencesInput, fn func(*codecommit.GetDifferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetDifferencesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetDifferencesOutput{}
	fnCacher := func(out *codecommit.GetDifferencesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetDifferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetDifferencesWithContext(ctx context.Context, input *codecommit.GetDifferencesInput, opts ...request.Option) (*codecommit.GetDifferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetDifferencesOutput), nil
	}
	out, err := c.CodeCommitAPI.GetDifferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetFile(input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetFileOutput), nil
	}
	out, err := c.CodeCommitAPI.GetFile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetFileWithContext(ctx context.Context, input *codecommit.GetFileInput, opts ...request.Option) (*codecommit.GetFileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetFileOutput), nil
	}
	out, err := c.CodeCommitAPI.GetFileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetFolder(input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetFolderOutput), nil
	}
	out, err := c.CodeCommitAPI.GetFolder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetFolderWithContext(ctx context.Context, input *codecommit.GetFolderInput, opts ...request.Option) (*codecommit.GetFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetFolderOutput), nil
	}
	out, err := c.CodeCommitAPI.GetFolderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetMergeCommit(input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetMergeCommitOutput), nil
	}
	out, err := c.CodeCommitAPI.GetMergeCommit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetMergeCommitWithContext(ctx context.Context, input *codecommit.GetMergeCommitInput, opts ...request.Option) (*codecommit.GetMergeCommitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetMergeCommitOutput), nil
	}
	out, err := c.CodeCommitAPI.GetMergeCommitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetMergeConflicts(input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetMergeConflictsOutput), nil
	}
	out, err := c.CodeCommitAPI.GetMergeConflicts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetMergeConflictsPages(input *codecommit.GetMergeConflictsInput, fn func(*codecommit.GetMergeConflictsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetMergeConflictsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetMergeConflictsOutput{}
	fnCacher := func(out *codecommit.GetMergeConflictsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetMergeConflictsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetMergeConflictsPagesWithContext(ctx context.Context, input *codecommit.GetMergeConflictsInput, fn func(*codecommit.GetMergeConflictsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.GetMergeConflictsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.GetMergeConflictsOutput{}
	fnCacher := func(out *codecommit.GetMergeConflictsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.GetMergeConflictsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) GetMergeConflictsWithContext(ctx context.Context, input *codecommit.GetMergeConflictsInput, opts ...request.Option) (*codecommit.GetMergeConflictsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetMergeConflictsOutput), nil
	}
	out, err := c.CodeCommitAPI.GetMergeConflictsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetMergeOptions(input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetMergeOptionsOutput), nil
	}
	out, err := c.CodeCommitAPI.GetMergeOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetMergeOptionsWithContext(ctx context.Context, input *codecommit.GetMergeOptionsInput, opts ...request.Option) (*codecommit.GetMergeOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetMergeOptionsOutput), nil
	}
	out, err := c.CodeCommitAPI.GetMergeOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetPullRequestApprovalStates(input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetPullRequestApprovalStatesOutput), nil
	}
	out, err := c.CodeCommitAPI.GetPullRequestApprovalStates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetPullRequestApprovalStatesWithContext(ctx context.Context, input *codecommit.GetPullRequestApprovalStatesInput, opts ...request.Option) (*codecommit.GetPullRequestApprovalStatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetPullRequestApprovalStatesOutput), nil
	}
	out, err := c.CodeCommitAPI.GetPullRequestApprovalStatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetPullRequestOverrideState(input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetPullRequestOverrideStateOutput), nil
	}
	out, err := c.CodeCommitAPI.GetPullRequestOverrideState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetPullRequestOverrideStateWithContext(ctx context.Context, input *codecommit.GetPullRequestOverrideStateInput, opts ...request.Option) (*codecommit.GetPullRequestOverrideStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetPullRequestOverrideStateOutput), nil
	}
	out, err := c.CodeCommitAPI.GetPullRequestOverrideStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetPullRequestWithContext(ctx context.Context, input *codecommit.GetPullRequestInput, opts ...request.Option) (*codecommit.GetPullRequestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetPullRequestOutput), nil
	}
	out, err := c.CodeCommitAPI.GetPullRequestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetRepository(input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetRepositoryOutput), nil
	}
	out, err := c.CodeCommitAPI.GetRepository(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetRepositoryTriggers(input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetRepositoryTriggersOutput), nil
	}
	out, err := c.CodeCommitAPI.GetRepositoryTriggers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetRepositoryTriggersWithContext(ctx context.Context, input *codecommit.GetRepositoryTriggersInput, opts ...request.Option) (*codecommit.GetRepositoryTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetRepositoryTriggersOutput), nil
	}
	out, err := c.CodeCommitAPI.GetRepositoryTriggersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) GetRepositoryWithContext(ctx context.Context, input *codecommit.GetRepositoryInput, opts ...request.Option) (*codecommit.GetRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.GetRepositoryOutput), nil
	}
	out, err := c.CodeCommitAPI.GetRepositoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListApprovalRuleTemplates(input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListApprovalRuleTemplatesOutput), nil
	}
	out, err := c.CodeCommitAPI.ListApprovalRuleTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListApprovalRuleTemplatesPages(input *codecommit.ListApprovalRuleTemplatesInput, fn func(*codecommit.ListApprovalRuleTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListApprovalRuleTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListApprovalRuleTemplatesOutput{}
	fnCacher := func(out *codecommit.ListApprovalRuleTemplatesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListApprovalRuleTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListApprovalRuleTemplatesPagesWithContext(ctx context.Context, input *codecommit.ListApprovalRuleTemplatesInput, fn func(*codecommit.ListApprovalRuleTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListApprovalRuleTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListApprovalRuleTemplatesOutput{}
	fnCacher := func(out *codecommit.ListApprovalRuleTemplatesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListApprovalRuleTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListApprovalRuleTemplatesWithContext(ctx context.Context, input *codecommit.ListApprovalRuleTemplatesInput, opts ...request.Option) (*codecommit.ListApprovalRuleTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListApprovalRuleTemplatesOutput), nil
	}
	out, err := c.CodeCommitAPI.ListApprovalRuleTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListAssociatedApprovalRuleTemplatesForRepository(input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput), nil
	}
	out, err := c.CodeCommitAPI.ListAssociatedApprovalRuleTemplatesForRepository(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListAssociatedApprovalRuleTemplatesForRepositoryPages(input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput, fn func(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput{}
	fnCacher := func(out *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListAssociatedApprovalRuleTemplatesForRepositoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListAssociatedApprovalRuleTemplatesForRepositoryPagesWithContext(ctx context.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput, fn func(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput{}
	fnCacher := func(out *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListAssociatedApprovalRuleTemplatesForRepositoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListAssociatedApprovalRuleTemplatesForRepositoryWithContext(ctx context.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput, opts ...request.Option) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput), nil
	}
	out, err := c.CodeCommitAPI.ListAssociatedApprovalRuleTemplatesForRepositoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListBranches(input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListBranchesOutput), nil
	}
	out, err := c.CodeCommitAPI.ListBranches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListBranchesPages(input *codecommit.ListBranchesInput, fn func(*codecommit.ListBranchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListBranchesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListBranchesOutput{}
	fnCacher := func(out *codecommit.ListBranchesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListBranchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListBranchesPagesWithContext(ctx context.Context, input *codecommit.ListBranchesInput, fn func(*codecommit.ListBranchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListBranchesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListBranchesOutput{}
	fnCacher := func(out *codecommit.ListBranchesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListBranchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListBranchesWithContext(ctx context.Context, input *codecommit.ListBranchesInput, opts ...request.Option) (*codecommit.ListBranchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListBranchesOutput), nil
	}
	out, err := c.CodeCommitAPI.ListBranchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListPullRequests(input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListPullRequestsOutput), nil
	}
	out, err := c.CodeCommitAPI.ListPullRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListPullRequestsPages(input *codecommit.ListPullRequestsInput, fn func(*codecommit.ListPullRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListPullRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListPullRequestsOutput{}
	fnCacher := func(out *codecommit.ListPullRequestsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListPullRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListPullRequestsPagesWithContext(ctx context.Context, input *codecommit.ListPullRequestsInput, fn func(*codecommit.ListPullRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListPullRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListPullRequestsOutput{}
	fnCacher := func(out *codecommit.ListPullRequestsOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListPullRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListPullRequestsWithContext(ctx context.Context, input *codecommit.ListPullRequestsInput, opts ...request.Option) (*codecommit.ListPullRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListPullRequestsOutput), nil
	}
	out, err := c.CodeCommitAPI.ListPullRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListRepositories(input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.ListRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListRepositoriesForApprovalRuleTemplate(input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput), nil
	}
	out, err := c.CodeCommitAPI.ListRepositoriesForApprovalRuleTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListRepositoriesForApprovalRuleTemplatePages(input *codecommit.ListRepositoriesForApprovalRuleTemplateInput, fn func(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListRepositoriesForApprovalRuleTemplateOutput{}
	fnCacher := func(out *codecommit.ListRepositoriesForApprovalRuleTemplateOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListRepositoriesForApprovalRuleTemplatePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListRepositoriesForApprovalRuleTemplatePagesWithContext(ctx context.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput, fn func(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListRepositoriesForApprovalRuleTemplateOutput{}
	fnCacher := func(out *codecommit.ListRepositoriesForApprovalRuleTemplateOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListRepositoriesForApprovalRuleTemplatePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListRepositoriesForApprovalRuleTemplateWithContext(ctx context.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput, opts ...request.Option) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListRepositoriesForApprovalRuleTemplateOutput), nil
	}
	out, err := c.CodeCommitAPI.ListRepositoriesForApprovalRuleTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListRepositoriesPages(input *codecommit.ListRepositoriesInput, fn func(*codecommit.ListRepositoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListRepositoriesOutput{}
	fnCacher := func(out *codecommit.ListRepositoriesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListRepositoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListRepositoriesPagesWithContext(ctx context.Context, input *codecommit.ListRepositoriesInput, fn func(*codecommit.ListRepositoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codecommit.ListRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &codecommit.ListRepositoriesOutput{}
	fnCacher := func(out *codecommit.ListRepositoriesOutput, more bool) bool {
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
	if err := c.CodeCommitAPI.ListRepositoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeCommit) ListRepositoriesWithContext(ctx context.Context, input *codecommit.ListRepositoriesInput, opts ...request.Option) (*codecommit.ListRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListRepositoriesOutput), nil
	}
	out, err := c.CodeCommitAPI.ListRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListTagsForResource(input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeCommitAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeCommit) ListTagsForResourceWithContext(ctx context.Context, input *codecommit.ListTagsForResourceInput, opts ...request.Option) (*codecommit.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codecommit.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeCommitAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

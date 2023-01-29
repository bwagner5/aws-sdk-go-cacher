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
package cleanroomscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cleanrooms"
	"github.com/aws/aws-sdk-go/service/cleanrooms/cleanroomsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CleanRooms struct {
	cleanroomsiface.CleanRoomsAPI
	cache *cache.Cache
}

func New(cleanroomsapi cleanroomsiface.CleanRoomsAPI) *CleanRooms {
	return &CleanRooms{
		CleanRoomsAPI: cleanroomsapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CleanRooms) BatchGetSchema(input *cleanrooms.BatchGetSchemaInput) (*cleanrooms.BatchGetSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.BatchGetSchemaOutput), nil
	}
	out, err := c.CleanRoomsAPI.BatchGetSchema(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) BatchGetSchemaWithContext(ctx context.Context, input *cleanrooms.BatchGetSchemaInput, opts ...request.Option) (*cleanrooms.BatchGetSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.BatchGetSchemaOutput), nil
	}
	out, err := c.CleanRoomsAPI.BatchGetSchemaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetCollaboration(input *cleanrooms.GetCollaborationInput) (*cleanrooms.GetCollaborationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetCollaborationOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetCollaboration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetCollaborationWithContext(ctx context.Context, input *cleanrooms.GetCollaborationInput, opts ...request.Option) (*cleanrooms.GetCollaborationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetCollaborationOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetCollaborationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetConfiguredTable(input *cleanrooms.GetConfiguredTableInput) (*cleanrooms.GetConfiguredTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetConfiguredTableOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetConfiguredTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetConfiguredTableAnalysisRule(input *cleanrooms.GetConfiguredTableAnalysisRuleInput) (*cleanrooms.GetConfiguredTableAnalysisRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetConfiguredTableAnalysisRuleOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetConfiguredTableAnalysisRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetConfiguredTableAnalysisRuleWithContext(ctx context.Context, input *cleanrooms.GetConfiguredTableAnalysisRuleInput, opts ...request.Option) (*cleanrooms.GetConfiguredTableAnalysisRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetConfiguredTableAnalysisRuleOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetConfiguredTableAnalysisRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetConfiguredTableAssociation(input *cleanrooms.GetConfiguredTableAssociationInput) (*cleanrooms.GetConfiguredTableAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetConfiguredTableAssociationOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetConfiguredTableAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetConfiguredTableAssociationWithContext(ctx context.Context, input *cleanrooms.GetConfiguredTableAssociationInput, opts ...request.Option) (*cleanrooms.GetConfiguredTableAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetConfiguredTableAssociationOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetConfiguredTableAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetConfiguredTableWithContext(ctx context.Context, input *cleanrooms.GetConfiguredTableInput, opts ...request.Option) (*cleanrooms.GetConfiguredTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetConfiguredTableOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetConfiguredTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetMembership(input *cleanrooms.GetMembershipInput) (*cleanrooms.GetMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetMembershipOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetMembershipWithContext(ctx context.Context, input *cleanrooms.GetMembershipInput, opts ...request.Option) (*cleanrooms.GetMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetMembershipOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetProtectedQuery(input *cleanrooms.GetProtectedQueryInput) (*cleanrooms.GetProtectedQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetProtectedQueryOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetProtectedQuery(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetProtectedQueryWithContext(ctx context.Context, input *cleanrooms.GetProtectedQueryInput, opts ...request.Option) (*cleanrooms.GetProtectedQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetProtectedQueryOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetProtectedQueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetSchema(input *cleanrooms.GetSchemaInput) (*cleanrooms.GetSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetSchemaOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetSchema(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetSchemaAnalysisRule(input *cleanrooms.GetSchemaAnalysisRuleInput) (*cleanrooms.GetSchemaAnalysisRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetSchemaAnalysisRuleOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetSchemaAnalysisRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetSchemaAnalysisRuleWithContext(ctx context.Context, input *cleanrooms.GetSchemaAnalysisRuleInput, opts ...request.Option) (*cleanrooms.GetSchemaAnalysisRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetSchemaAnalysisRuleOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetSchemaAnalysisRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) GetSchemaWithContext(ctx context.Context, input *cleanrooms.GetSchemaInput, opts ...request.Option) (*cleanrooms.GetSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.GetSchemaOutput), nil
	}
	out, err := c.CleanRoomsAPI.GetSchemaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListCollaborations(input *cleanrooms.ListCollaborationsInput) (*cleanrooms.ListCollaborationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListCollaborationsOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListCollaborations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListCollaborationsPages(input *cleanrooms.ListCollaborationsInput, fn func(*cleanrooms.ListCollaborationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListCollaborationsOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListCollaborationsOutput{}
	fnCacher := func(out *cleanrooms.ListCollaborationsOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListCollaborationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListCollaborationsPagesWithContext(ctx context.Context, input *cleanrooms.ListCollaborationsInput, fn func(*cleanrooms.ListCollaborationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListCollaborationsOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListCollaborationsOutput{}
	fnCacher := func(out *cleanrooms.ListCollaborationsOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListCollaborationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListCollaborationsWithContext(ctx context.Context, input *cleanrooms.ListCollaborationsInput, opts ...request.Option) (*cleanrooms.ListCollaborationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListCollaborationsOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListCollaborationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListConfiguredTableAssociations(input *cleanrooms.ListConfiguredTableAssociationsInput) (*cleanrooms.ListConfiguredTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListConfiguredTableAssociationsOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListConfiguredTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListConfiguredTableAssociationsPages(input *cleanrooms.ListConfiguredTableAssociationsInput, fn func(*cleanrooms.ListConfiguredTableAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListConfiguredTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListConfiguredTableAssociationsOutput{}
	fnCacher := func(out *cleanrooms.ListConfiguredTableAssociationsOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListConfiguredTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListConfiguredTableAssociationsPagesWithContext(ctx context.Context, input *cleanrooms.ListConfiguredTableAssociationsInput, fn func(*cleanrooms.ListConfiguredTableAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListConfiguredTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListConfiguredTableAssociationsOutput{}
	fnCacher := func(out *cleanrooms.ListConfiguredTableAssociationsOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListConfiguredTableAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListConfiguredTableAssociationsWithContext(ctx context.Context, input *cleanrooms.ListConfiguredTableAssociationsInput, opts ...request.Option) (*cleanrooms.ListConfiguredTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListConfiguredTableAssociationsOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListConfiguredTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListConfiguredTables(input *cleanrooms.ListConfiguredTablesInput) (*cleanrooms.ListConfiguredTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListConfiguredTablesOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListConfiguredTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListConfiguredTablesPages(input *cleanrooms.ListConfiguredTablesInput, fn func(*cleanrooms.ListConfiguredTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListConfiguredTablesOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListConfiguredTablesOutput{}
	fnCacher := func(out *cleanrooms.ListConfiguredTablesOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListConfiguredTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListConfiguredTablesPagesWithContext(ctx context.Context, input *cleanrooms.ListConfiguredTablesInput, fn func(*cleanrooms.ListConfiguredTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListConfiguredTablesOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListConfiguredTablesOutput{}
	fnCacher := func(out *cleanrooms.ListConfiguredTablesOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListConfiguredTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListConfiguredTablesWithContext(ctx context.Context, input *cleanrooms.ListConfiguredTablesInput, opts ...request.Option) (*cleanrooms.ListConfiguredTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListConfiguredTablesOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListConfiguredTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListMembers(input *cleanrooms.ListMembersInput) (*cleanrooms.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListMembersOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListMembersPages(input *cleanrooms.ListMembersInput, fn func(*cleanrooms.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListMembersOutput{}
	fnCacher := func(out *cleanrooms.ListMembersOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListMembersPagesWithContext(ctx context.Context, input *cleanrooms.ListMembersInput, fn func(*cleanrooms.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListMembersOutput{}
	fnCacher := func(out *cleanrooms.ListMembersOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListMembersWithContext(ctx context.Context, input *cleanrooms.ListMembersInput, opts ...request.Option) (*cleanrooms.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListMembersOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListMemberships(input *cleanrooms.ListMembershipsInput) (*cleanrooms.ListMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListMembershipsOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListMembershipsPages(input *cleanrooms.ListMembershipsInput, fn func(*cleanrooms.ListMembershipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListMembershipsOutput{}
	fnCacher := func(out *cleanrooms.ListMembershipsOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListMembershipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListMembershipsPagesWithContext(ctx context.Context, input *cleanrooms.ListMembershipsInput, fn func(*cleanrooms.ListMembershipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListMembershipsOutput{}
	fnCacher := func(out *cleanrooms.ListMembershipsOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListMembershipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListMembershipsWithContext(ctx context.Context, input *cleanrooms.ListMembershipsInput, opts ...request.Option) (*cleanrooms.ListMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListMembershipsOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListProtectedQueries(input *cleanrooms.ListProtectedQueriesInput) (*cleanrooms.ListProtectedQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListProtectedQueriesOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListProtectedQueries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListProtectedQueriesPages(input *cleanrooms.ListProtectedQueriesInput, fn func(*cleanrooms.ListProtectedQueriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListProtectedQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListProtectedQueriesOutput{}
	fnCacher := func(out *cleanrooms.ListProtectedQueriesOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListProtectedQueriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListProtectedQueriesPagesWithContext(ctx context.Context, input *cleanrooms.ListProtectedQueriesInput, fn func(*cleanrooms.ListProtectedQueriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListProtectedQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListProtectedQueriesOutput{}
	fnCacher := func(out *cleanrooms.ListProtectedQueriesOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListProtectedQueriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListProtectedQueriesWithContext(ctx context.Context, input *cleanrooms.ListProtectedQueriesInput, opts ...request.Option) (*cleanrooms.ListProtectedQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListProtectedQueriesOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListProtectedQueriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListSchemas(input *cleanrooms.ListSchemasInput) (*cleanrooms.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListSchemasOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListSchemas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CleanRooms) ListSchemasPages(input *cleanrooms.ListSchemasInput, fn func(*cleanrooms.ListSchemasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListSchemasOutput{}
	fnCacher := func(out *cleanrooms.ListSchemasOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListSchemasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListSchemasPagesWithContext(ctx context.Context, input *cleanrooms.ListSchemasInput, fn func(*cleanrooms.ListSchemasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cleanrooms.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &cleanrooms.ListSchemasOutput{}
	fnCacher := func(out *cleanrooms.ListSchemasOutput, more bool) bool {
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
	if err := c.CleanRoomsAPI.ListSchemasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CleanRooms) ListSchemasWithContext(ctx context.Context, input *cleanrooms.ListSchemasInput, opts ...request.Option) (*cleanrooms.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cleanrooms.ListSchemasOutput), nil
	}
	out, err := c.CleanRoomsAPI.ListSchemasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

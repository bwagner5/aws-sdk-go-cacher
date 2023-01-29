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
package budgetscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/budgets/budgetsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Budgets struct {
	budgetsiface.BudgetsAPI
	cache *cache.Cache
}

func New(budgetsapi budgetsiface.BudgetsAPI) *Budgets {
	return &Budgets{
		BudgetsAPI: budgetsapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Budgets) DescribeBudget(input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetAction(input *budgets.DescribeBudgetActionInput) (*budgets.DescribeBudgetActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionHistories(input *budgets.DescribeBudgetActionHistoriesInput) (*budgets.DescribeBudgetActionHistoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionHistoriesOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionHistories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionHistoriesPages(input *budgets.DescribeBudgetActionHistoriesInput, fn func(*budgets.DescribeBudgetActionHistoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetActionHistoriesOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetActionHistoriesOutput{}
	fnCacher := func(out *budgets.DescribeBudgetActionHistoriesOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetActionHistoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetActionHistoriesPagesWithContext(ctx context.Context, input *budgets.DescribeBudgetActionHistoriesInput, fn func(*budgets.DescribeBudgetActionHistoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetActionHistoriesOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetActionHistoriesOutput{}
	fnCacher := func(out *budgets.DescribeBudgetActionHistoriesOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetActionHistoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetActionHistoriesWithContext(ctx context.Context, input *budgets.DescribeBudgetActionHistoriesInput, opts ...request.Option) (*budgets.DescribeBudgetActionHistoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionHistoriesOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionHistoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionWithContext(ctx context.Context, input *budgets.DescribeBudgetActionInput, opts ...request.Option) (*budgets.DescribeBudgetActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionsForAccount(input *budgets.DescribeBudgetActionsForAccountInput) (*budgets.DescribeBudgetActionsForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionsForAccountOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionsForAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionsForAccountPages(input *budgets.DescribeBudgetActionsForAccountInput, fn func(*budgets.DescribeBudgetActionsForAccountOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetActionsForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetActionsForAccountOutput{}
	fnCacher := func(out *budgets.DescribeBudgetActionsForAccountOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetActionsForAccountPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetActionsForAccountPagesWithContext(ctx context.Context, input *budgets.DescribeBudgetActionsForAccountInput, fn func(*budgets.DescribeBudgetActionsForAccountOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetActionsForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetActionsForAccountOutput{}
	fnCacher := func(out *budgets.DescribeBudgetActionsForAccountOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetActionsForAccountPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetActionsForAccountWithContext(ctx context.Context, input *budgets.DescribeBudgetActionsForAccountInput, opts ...request.Option) (*budgets.DescribeBudgetActionsForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionsForAccountOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionsForAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionsForBudget(input *budgets.DescribeBudgetActionsForBudgetInput) (*budgets.DescribeBudgetActionsForBudgetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionsForBudgetOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionsForBudget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetActionsForBudgetPages(input *budgets.DescribeBudgetActionsForBudgetInput, fn func(*budgets.DescribeBudgetActionsForBudgetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetActionsForBudgetOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetActionsForBudgetOutput{}
	fnCacher := func(out *budgets.DescribeBudgetActionsForBudgetOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetActionsForBudgetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetActionsForBudgetPagesWithContext(ctx context.Context, input *budgets.DescribeBudgetActionsForBudgetInput, fn func(*budgets.DescribeBudgetActionsForBudgetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetActionsForBudgetOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetActionsForBudgetOutput{}
	fnCacher := func(out *budgets.DescribeBudgetActionsForBudgetOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetActionsForBudgetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetActionsForBudgetWithContext(ctx context.Context, input *budgets.DescribeBudgetActionsForBudgetInput, opts ...request.Option) (*budgets.DescribeBudgetActionsForBudgetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetActionsForBudgetOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetActionsForBudgetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetNotificationsForAccount(input *budgets.DescribeBudgetNotificationsForAccountInput) (*budgets.DescribeBudgetNotificationsForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetNotificationsForAccountOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetNotificationsForAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetNotificationsForAccountPages(input *budgets.DescribeBudgetNotificationsForAccountInput, fn func(*budgets.DescribeBudgetNotificationsForAccountOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetNotificationsForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetNotificationsForAccountOutput{}
	fnCacher := func(out *budgets.DescribeBudgetNotificationsForAccountOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetNotificationsForAccountPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetNotificationsForAccountPagesWithContext(ctx context.Context, input *budgets.DescribeBudgetNotificationsForAccountInput, fn func(*budgets.DescribeBudgetNotificationsForAccountOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetNotificationsForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetNotificationsForAccountOutput{}
	fnCacher := func(out *budgets.DescribeBudgetNotificationsForAccountOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetNotificationsForAccountPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetNotificationsForAccountWithContext(ctx context.Context, input *budgets.DescribeBudgetNotificationsForAccountInput, opts ...request.Option) (*budgets.DescribeBudgetNotificationsForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetNotificationsForAccountOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetNotificationsForAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetPerformanceHistory(input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetPerformanceHistoryOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetPerformanceHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetPerformanceHistoryPages(input *budgets.DescribeBudgetPerformanceHistoryInput, fn func(*budgets.DescribeBudgetPerformanceHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetPerformanceHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetPerformanceHistoryOutput{}
	fnCacher := func(out *budgets.DescribeBudgetPerformanceHistoryOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetPerformanceHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetPerformanceHistoryPagesWithContext(ctx context.Context, input *budgets.DescribeBudgetPerformanceHistoryInput, fn func(*budgets.DescribeBudgetPerformanceHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetPerformanceHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetPerformanceHistoryOutput{}
	fnCacher := func(out *budgets.DescribeBudgetPerformanceHistoryOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetPerformanceHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetPerformanceHistoryWithContext(ctx context.Context, input *budgets.DescribeBudgetPerformanceHistoryInput, opts ...request.Option) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetPerformanceHistoryOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetPerformanceHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetWithContext(ctx context.Context, input *budgets.DescribeBudgetInput, opts ...request.Option) (*budgets.DescribeBudgetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgets(input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetsOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeBudgetsPages(input *budgets.DescribeBudgetsInput, fn func(*budgets.DescribeBudgetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetsOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetsOutput{}
	fnCacher := func(out *budgets.DescribeBudgetsOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetsPagesWithContext(ctx context.Context, input *budgets.DescribeBudgetsInput, fn func(*budgets.DescribeBudgetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeBudgetsOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeBudgetsOutput{}
	fnCacher := func(out *budgets.DescribeBudgetsOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeBudgetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeBudgetsWithContext(ctx context.Context, input *budgets.DescribeBudgetsInput, opts ...request.Option) (*budgets.DescribeBudgetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeBudgetsOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeBudgetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeNotificationsForBudget(input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeNotificationsForBudgetOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeNotificationsForBudget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeNotificationsForBudgetPages(input *budgets.DescribeNotificationsForBudgetInput, fn func(*budgets.DescribeNotificationsForBudgetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeNotificationsForBudgetOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeNotificationsForBudgetOutput{}
	fnCacher := func(out *budgets.DescribeNotificationsForBudgetOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeNotificationsForBudgetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeNotificationsForBudgetPagesWithContext(ctx context.Context, input *budgets.DescribeNotificationsForBudgetInput, fn func(*budgets.DescribeNotificationsForBudgetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeNotificationsForBudgetOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeNotificationsForBudgetOutput{}
	fnCacher := func(out *budgets.DescribeNotificationsForBudgetOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeNotificationsForBudgetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeNotificationsForBudgetWithContext(ctx context.Context, input *budgets.DescribeNotificationsForBudgetInput, opts ...request.Option) (*budgets.DescribeNotificationsForBudgetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeNotificationsForBudgetOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeNotificationsForBudgetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeSubscribersForNotification(input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeSubscribersForNotificationOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeSubscribersForNotification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Budgets) DescribeSubscribersForNotificationPages(input *budgets.DescribeSubscribersForNotificationInput, fn func(*budgets.DescribeSubscribersForNotificationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeSubscribersForNotificationOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeSubscribersForNotificationOutput{}
	fnCacher := func(out *budgets.DescribeSubscribersForNotificationOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeSubscribersForNotificationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeSubscribersForNotificationPagesWithContext(ctx context.Context, input *budgets.DescribeSubscribersForNotificationInput, fn func(*budgets.DescribeSubscribersForNotificationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*budgets.DescribeSubscribersForNotificationOutput), false)
		return nil
	}
	cachable := true
	output := &budgets.DescribeSubscribersForNotificationOutput{}
	fnCacher := func(out *budgets.DescribeSubscribersForNotificationOutput, more bool) bool {
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
	if err := c.BudgetsAPI.DescribeSubscribersForNotificationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Budgets) DescribeSubscribersForNotificationWithContext(ctx context.Context, input *budgets.DescribeSubscribersForNotificationInput, opts ...request.Option) (*budgets.DescribeSubscribersForNotificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*budgets.DescribeSubscribersForNotificationOutput), nil
	}
	out, err := c.BudgetsAPI.DescribeSubscribersForNotificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

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
package servicecatalogcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ServiceCatalog struct {
	servicecatalogiface.ServiceCatalogAPI
	cache *cache.Cache
}

func New(servicecatalogapi servicecatalogiface.ServiceCatalogAPI) *ServiceCatalog {
	return &ServiceCatalog{
		ServiceCatalogAPI: servicecatalogapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ServiceCatalog) BatchAssociateServiceActionWithProvisioningArtifact(input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.BatchAssociateServiceActionWithProvisioningArtifact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) BatchAssociateServiceActionWithProvisioningArtifactWithContext(ctx context.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput, opts ...request.Option) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.BatchAssociateServiceActionWithProvisioningArtifactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) BatchDisassociateServiceActionFromProvisioningArtifact(input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.BatchDisassociateServiceActionFromProvisioningArtifact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) BatchDisassociateServiceActionFromProvisioningArtifactWithContext(ctx context.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput, opts ...request.Option) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.BatchDisassociateServiceActionFromProvisioningArtifactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeConstraint(input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeConstraintOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeConstraint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeConstraintWithContext(ctx context.Context, input *servicecatalog.DescribeConstraintInput, opts ...request.Option) (*servicecatalog.DescribeConstraintOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeConstraintOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeConstraintWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeCopyProductStatus(input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeCopyProductStatusOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeCopyProductStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeCopyProductStatusWithContext(ctx context.Context, input *servicecatalog.DescribeCopyProductStatusInput, opts ...request.Option) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeCopyProductStatusOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeCopyProductStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribePortfolio(input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribePortfolioOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribePortfolio(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribePortfolioShareStatus(input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribePortfolioShareStatusOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribePortfolioShareStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribePortfolioShareStatusWithContext(ctx context.Context, input *servicecatalog.DescribePortfolioShareStatusInput, opts ...request.Option) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribePortfolioShareStatusOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribePortfolioShareStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribePortfolioShares(input *servicecatalog.DescribePortfolioSharesInput) (*servicecatalog.DescribePortfolioSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribePortfolioSharesOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribePortfolioShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribePortfolioSharesPages(input *servicecatalog.DescribePortfolioSharesInput, fn func(*servicecatalog.DescribePortfolioSharesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.DescribePortfolioSharesOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.DescribePortfolioSharesOutput{}
	fnCacher := func(out *servicecatalog.DescribePortfolioSharesOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.DescribePortfolioSharesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) DescribePortfolioSharesPagesWithContext(ctx context.Context, input *servicecatalog.DescribePortfolioSharesInput, fn func(*servicecatalog.DescribePortfolioSharesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.DescribePortfolioSharesOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.DescribePortfolioSharesOutput{}
	fnCacher := func(out *servicecatalog.DescribePortfolioSharesOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.DescribePortfolioSharesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) DescribePortfolioSharesWithContext(ctx context.Context, input *servicecatalog.DescribePortfolioSharesInput, opts ...request.Option) (*servicecatalog.DescribePortfolioSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribePortfolioSharesOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribePortfolioSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribePortfolioWithContext(ctx context.Context, input *servicecatalog.DescribePortfolioInput, opts ...request.Option) (*servicecatalog.DescribePortfolioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribePortfolioOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribePortfolioWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProduct(input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProduct(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProductAsAdmin(input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProductAsAdminOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProductAsAdmin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProductAsAdminWithContext(ctx context.Context, input *servicecatalog.DescribeProductAsAdminInput, opts ...request.Option) (*servicecatalog.DescribeProductAsAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProductAsAdminOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProductAsAdminWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProductView(input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProductViewOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProductView(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProductViewWithContext(ctx context.Context, input *servicecatalog.DescribeProductViewInput, opts ...request.Option) (*servicecatalog.DescribeProductViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProductViewOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProductViewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProductWithContext(ctx context.Context, input *servicecatalog.DescribeProductInput, opts ...request.Option) (*servicecatalog.DescribeProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProductWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisionedProduct(input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisionedProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisionedProduct(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisionedProductPlan(input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisionedProductPlanOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisionedProductPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisionedProductPlanWithContext(ctx context.Context, input *servicecatalog.DescribeProvisionedProductPlanInput, opts ...request.Option) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisionedProductPlanOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisionedProductPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisionedProductWithContext(ctx context.Context, input *servicecatalog.DescribeProvisionedProductInput, opts ...request.Option) (*servicecatalog.DescribeProvisionedProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisionedProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisionedProductWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisioningArtifact(input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisioningArtifact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisioningArtifactWithContext(ctx context.Context, input *servicecatalog.DescribeProvisioningArtifactInput, opts ...request.Option) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisioningArtifactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisioningParameters(input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisioningParametersOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisioningParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeProvisioningParametersWithContext(ctx context.Context, input *servicecatalog.DescribeProvisioningParametersInput, opts ...request.Option) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeProvisioningParametersOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeProvisioningParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeRecord(input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeRecordOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeRecord(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeRecordWithContext(ctx context.Context, input *servicecatalog.DescribeRecordInput, opts ...request.Option) (*servicecatalog.DescribeRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeRecordOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeRecordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeServiceAction(input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeServiceActionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeServiceAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeServiceActionExecutionParameters(input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeServiceActionExecutionParametersOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeServiceActionExecutionParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeServiceActionExecutionParametersWithContext(ctx context.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput, opts ...request.Option) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeServiceActionExecutionParametersOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeServiceActionExecutionParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeServiceActionWithContext(ctx context.Context, input *servicecatalog.DescribeServiceActionInput, opts ...request.Option) (*servicecatalog.DescribeServiceActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeServiceActionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeServiceActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeTagOption(input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeTagOptionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeTagOption(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) DescribeTagOptionWithContext(ctx context.Context, input *servicecatalog.DescribeTagOptionInput, opts ...request.Option) (*servicecatalog.DescribeTagOptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.DescribeTagOptionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.DescribeTagOptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) GetAWSOrganizationsAccessStatus(input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.GetAWSOrganizationsAccessStatusOutput), nil
	}
	out, err := c.ServiceCatalogAPI.GetAWSOrganizationsAccessStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) GetAWSOrganizationsAccessStatusWithContext(ctx context.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput, opts ...request.Option) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.GetAWSOrganizationsAccessStatusOutput), nil
	}
	out, err := c.ServiceCatalogAPI.GetAWSOrganizationsAccessStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) GetProvisionedProductOutputs(input *servicecatalog.GetProvisionedProductOutputsInput) (*servicecatalog.GetProvisionedProductOutputsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.GetProvisionedProductOutputsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.GetProvisionedProductOutputs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) GetProvisionedProductOutputsPages(input *servicecatalog.GetProvisionedProductOutputsInput, fn func(*servicecatalog.GetProvisionedProductOutputsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.GetProvisionedProductOutputsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.GetProvisionedProductOutputsOutput{}
	fnCacher := func(out *servicecatalog.GetProvisionedProductOutputsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.GetProvisionedProductOutputsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) GetProvisionedProductOutputsPagesWithContext(ctx context.Context, input *servicecatalog.GetProvisionedProductOutputsInput, fn func(*servicecatalog.GetProvisionedProductOutputsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.GetProvisionedProductOutputsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.GetProvisionedProductOutputsOutput{}
	fnCacher := func(out *servicecatalog.GetProvisionedProductOutputsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.GetProvisionedProductOutputsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) GetProvisionedProductOutputsWithContext(ctx context.Context, input *servicecatalog.GetProvisionedProductOutputsInput, opts ...request.Option) (*servicecatalog.GetProvisionedProductOutputsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.GetProvisionedProductOutputsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.GetProvisionedProductOutputsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListAcceptedPortfolioShares(input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListAcceptedPortfolioSharesOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListAcceptedPortfolioShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListAcceptedPortfolioSharesPages(input *servicecatalog.ListAcceptedPortfolioSharesInput, fn func(*servicecatalog.ListAcceptedPortfolioSharesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListAcceptedPortfolioSharesOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListAcceptedPortfolioSharesOutput{}
	fnCacher := func(out *servicecatalog.ListAcceptedPortfolioSharesOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListAcceptedPortfolioSharesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListAcceptedPortfolioSharesPagesWithContext(ctx context.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput, fn func(*servicecatalog.ListAcceptedPortfolioSharesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListAcceptedPortfolioSharesOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListAcceptedPortfolioSharesOutput{}
	fnCacher := func(out *servicecatalog.ListAcceptedPortfolioSharesOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListAcceptedPortfolioSharesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListAcceptedPortfolioSharesWithContext(ctx context.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput, opts ...request.Option) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListAcceptedPortfolioSharesOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListAcceptedPortfolioSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListBudgetsForResource(input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListBudgetsForResourceOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListBudgetsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListBudgetsForResourcePages(input *servicecatalog.ListBudgetsForResourceInput, fn func(*servicecatalog.ListBudgetsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListBudgetsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListBudgetsForResourceOutput{}
	fnCacher := func(out *servicecatalog.ListBudgetsForResourceOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListBudgetsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListBudgetsForResourcePagesWithContext(ctx context.Context, input *servicecatalog.ListBudgetsForResourceInput, fn func(*servicecatalog.ListBudgetsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListBudgetsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListBudgetsForResourceOutput{}
	fnCacher := func(out *servicecatalog.ListBudgetsForResourceOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListBudgetsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListBudgetsForResourceWithContext(ctx context.Context, input *servicecatalog.ListBudgetsForResourceInput, opts ...request.Option) (*servicecatalog.ListBudgetsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListBudgetsForResourceOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListBudgetsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListConstraintsForPortfolio(input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListConstraintsForPortfolioOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListConstraintsForPortfolio(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListConstraintsForPortfolioPages(input *servicecatalog.ListConstraintsForPortfolioInput, fn func(*servicecatalog.ListConstraintsForPortfolioOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListConstraintsForPortfolioOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListConstraintsForPortfolioOutput{}
	fnCacher := func(out *servicecatalog.ListConstraintsForPortfolioOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListConstraintsForPortfolioPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListConstraintsForPortfolioPagesWithContext(ctx context.Context, input *servicecatalog.ListConstraintsForPortfolioInput, fn func(*servicecatalog.ListConstraintsForPortfolioOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListConstraintsForPortfolioOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListConstraintsForPortfolioOutput{}
	fnCacher := func(out *servicecatalog.ListConstraintsForPortfolioOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListConstraintsForPortfolioPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListConstraintsForPortfolioWithContext(ctx context.Context, input *servicecatalog.ListConstraintsForPortfolioInput, opts ...request.Option) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListConstraintsForPortfolioOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListConstraintsForPortfolioWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListLaunchPaths(input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListLaunchPathsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListLaunchPaths(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListLaunchPathsPages(input *servicecatalog.ListLaunchPathsInput, fn func(*servicecatalog.ListLaunchPathsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListLaunchPathsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListLaunchPathsOutput{}
	fnCacher := func(out *servicecatalog.ListLaunchPathsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListLaunchPathsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListLaunchPathsPagesWithContext(ctx context.Context, input *servicecatalog.ListLaunchPathsInput, fn func(*servicecatalog.ListLaunchPathsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListLaunchPathsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListLaunchPathsOutput{}
	fnCacher := func(out *servicecatalog.ListLaunchPathsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListLaunchPathsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListLaunchPathsWithContext(ctx context.Context, input *servicecatalog.ListLaunchPathsInput, opts ...request.Option) (*servicecatalog.ListLaunchPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListLaunchPathsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListLaunchPathsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListOrganizationPortfolioAccess(input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListOrganizationPortfolioAccessOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListOrganizationPortfolioAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListOrganizationPortfolioAccessPages(input *servicecatalog.ListOrganizationPortfolioAccessInput, fn func(*servicecatalog.ListOrganizationPortfolioAccessOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListOrganizationPortfolioAccessOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListOrganizationPortfolioAccessOutput{}
	fnCacher := func(out *servicecatalog.ListOrganizationPortfolioAccessOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListOrganizationPortfolioAccessPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListOrganizationPortfolioAccessPagesWithContext(ctx context.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput, fn func(*servicecatalog.ListOrganizationPortfolioAccessOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListOrganizationPortfolioAccessOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListOrganizationPortfolioAccessOutput{}
	fnCacher := func(out *servicecatalog.ListOrganizationPortfolioAccessOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListOrganizationPortfolioAccessPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListOrganizationPortfolioAccessWithContext(ctx context.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput, opts ...request.Option) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListOrganizationPortfolioAccessOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListOrganizationPortfolioAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPortfolioAccess(input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPortfolioAccessOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPortfolioAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPortfolioAccessPages(input *servicecatalog.ListPortfolioAccessInput, fn func(*servicecatalog.ListPortfolioAccessOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPortfolioAccessOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPortfolioAccessOutput{}
	fnCacher := func(out *servicecatalog.ListPortfolioAccessOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPortfolioAccessPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPortfolioAccessPagesWithContext(ctx context.Context, input *servicecatalog.ListPortfolioAccessInput, fn func(*servicecatalog.ListPortfolioAccessOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPortfolioAccessOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPortfolioAccessOutput{}
	fnCacher := func(out *servicecatalog.ListPortfolioAccessOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPortfolioAccessPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPortfolioAccessWithContext(ctx context.Context, input *servicecatalog.ListPortfolioAccessInput, opts ...request.Option) (*servicecatalog.ListPortfolioAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPortfolioAccessOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPortfolioAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPortfolios(input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPortfoliosOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPortfolios(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPortfoliosForProduct(input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPortfoliosForProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPortfoliosForProduct(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPortfoliosForProductPages(input *servicecatalog.ListPortfoliosForProductInput, fn func(*servicecatalog.ListPortfoliosForProductOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPortfoliosForProductOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPortfoliosForProductOutput{}
	fnCacher := func(out *servicecatalog.ListPortfoliosForProductOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPortfoliosForProductPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPortfoliosForProductPagesWithContext(ctx context.Context, input *servicecatalog.ListPortfoliosForProductInput, fn func(*servicecatalog.ListPortfoliosForProductOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPortfoliosForProductOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPortfoliosForProductOutput{}
	fnCacher := func(out *servicecatalog.ListPortfoliosForProductOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPortfoliosForProductPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPortfoliosForProductWithContext(ctx context.Context, input *servicecatalog.ListPortfoliosForProductInput, opts ...request.Option) (*servicecatalog.ListPortfoliosForProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPortfoliosForProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPortfoliosForProductWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPortfoliosPages(input *servicecatalog.ListPortfoliosInput, fn func(*servicecatalog.ListPortfoliosOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPortfoliosOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPortfoliosOutput{}
	fnCacher := func(out *servicecatalog.ListPortfoliosOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPortfoliosPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPortfoliosPagesWithContext(ctx context.Context, input *servicecatalog.ListPortfoliosInput, fn func(*servicecatalog.ListPortfoliosOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPortfoliosOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPortfoliosOutput{}
	fnCacher := func(out *servicecatalog.ListPortfoliosOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPortfoliosPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPortfoliosWithContext(ctx context.Context, input *servicecatalog.ListPortfoliosInput, opts ...request.Option) (*servicecatalog.ListPortfoliosOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPortfoliosOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPortfoliosWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPrincipalsForPortfolio(input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPrincipalsForPortfolioOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPrincipalsForPortfolio(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListPrincipalsForPortfolioPages(input *servicecatalog.ListPrincipalsForPortfolioInput, fn func(*servicecatalog.ListPrincipalsForPortfolioOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPrincipalsForPortfolioOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPrincipalsForPortfolioOutput{}
	fnCacher := func(out *servicecatalog.ListPrincipalsForPortfolioOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPrincipalsForPortfolioPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPrincipalsForPortfolioPagesWithContext(ctx context.Context, input *servicecatalog.ListPrincipalsForPortfolioInput, fn func(*servicecatalog.ListPrincipalsForPortfolioOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListPrincipalsForPortfolioOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListPrincipalsForPortfolioOutput{}
	fnCacher := func(out *servicecatalog.ListPrincipalsForPortfolioOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListPrincipalsForPortfolioPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListPrincipalsForPortfolioWithContext(ctx context.Context, input *servicecatalog.ListPrincipalsForPortfolioInput, opts ...request.Option) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListPrincipalsForPortfolioOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListPrincipalsForPortfolioWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListProvisionedProductPlans(input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListProvisionedProductPlansOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListProvisionedProductPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListProvisionedProductPlansWithContext(ctx context.Context, input *servicecatalog.ListProvisionedProductPlansInput, opts ...request.Option) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListProvisionedProductPlansOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListProvisionedProductPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListProvisioningArtifacts(input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListProvisioningArtifactsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListProvisioningArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListProvisioningArtifactsForServiceAction(input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListProvisioningArtifactsForServiceAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListProvisioningArtifactsForServiceActionPages(input *servicecatalog.ListProvisioningArtifactsForServiceActionInput, fn func(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListProvisioningArtifactsForServiceActionOutput{}
	fnCacher := func(out *servicecatalog.ListProvisioningArtifactsForServiceActionOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListProvisioningArtifactsForServiceActionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListProvisioningArtifactsForServiceActionPagesWithContext(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput, fn func(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListProvisioningArtifactsForServiceActionOutput{}
	fnCacher := func(out *servicecatalog.ListProvisioningArtifactsForServiceActionOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListProvisioningArtifactsForServiceActionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListProvisioningArtifactsForServiceActionWithContext(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput, opts ...request.Option) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListProvisioningArtifactsForServiceActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListProvisioningArtifactsWithContext(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsInput, opts ...request.Option) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListProvisioningArtifactsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListProvisioningArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListRecordHistory(input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListRecordHistoryOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListRecordHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListRecordHistoryWithContext(ctx context.Context, input *servicecatalog.ListRecordHistoryInput, opts ...request.Option) (*servicecatalog.ListRecordHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListRecordHistoryOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListRecordHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListResourcesForTagOption(input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListResourcesForTagOptionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListResourcesForTagOption(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListResourcesForTagOptionPages(input *servicecatalog.ListResourcesForTagOptionInput, fn func(*servicecatalog.ListResourcesForTagOptionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListResourcesForTagOptionOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListResourcesForTagOptionOutput{}
	fnCacher := func(out *servicecatalog.ListResourcesForTagOptionOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListResourcesForTagOptionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListResourcesForTagOptionPagesWithContext(ctx context.Context, input *servicecatalog.ListResourcesForTagOptionInput, fn func(*servicecatalog.ListResourcesForTagOptionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListResourcesForTagOptionOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListResourcesForTagOptionOutput{}
	fnCacher := func(out *servicecatalog.ListResourcesForTagOptionOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListResourcesForTagOptionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListResourcesForTagOptionWithContext(ctx context.Context, input *servicecatalog.ListResourcesForTagOptionInput, opts ...request.Option) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListResourcesForTagOptionOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListResourcesForTagOptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListServiceActions(input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListServiceActionsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListServiceActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListServiceActionsForProvisioningArtifact(input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListServiceActionsForProvisioningArtifact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListServiceActionsForProvisioningArtifactPages(input *servicecatalog.ListServiceActionsForProvisioningArtifactInput, fn func(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListServiceActionsForProvisioningArtifactOutput{}
	fnCacher := func(out *servicecatalog.ListServiceActionsForProvisioningArtifactOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListServiceActionsForProvisioningArtifactPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListServiceActionsForProvisioningArtifactPagesWithContext(ctx context.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput, fn func(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListServiceActionsForProvisioningArtifactOutput{}
	fnCacher := func(out *servicecatalog.ListServiceActionsForProvisioningArtifactOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListServiceActionsForProvisioningArtifactPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListServiceActionsForProvisioningArtifactWithContext(ctx context.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput, opts ...request.Option) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListServiceActionsForProvisioningArtifactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListServiceActionsPages(input *servicecatalog.ListServiceActionsInput, fn func(*servicecatalog.ListServiceActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListServiceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListServiceActionsOutput{}
	fnCacher := func(out *servicecatalog.ListServiceActionsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListServiceActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListServiceActionsPagesWithContext(ctx context.Context, input *servicecatalog.ListServiceActionsInput, fn func(*servicecatalog.ListServiceActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListServiceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListServiceActionsOutput{}
	fnCacher := func(out *servicecatalog.ListServiceActionsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListServiceActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListServiceActionsWithContext(ctx context.Context, input *servicecatalog.ListServiceActionsInput, opts ...request.Option) (*servicecatalog.ListServiceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListServiceActionsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListServiceActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListStackInstancesForProvisionedProduct(input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListStackInstancesForProvisionedProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListStackInstancesForProvisionedProduct(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListStackInstancesForProvisionedProductWithContext(ctx context.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput, opts ...request.Option) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListStackInstancesForProvisionedProductOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListStackInstancesForProvisionedProductWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListTagOptions(input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListTagOptionsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListTagOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ListTagOptionsPages(input *servicecatalog.ListTagOptionsInput, fn func(*servicecatalog.ListTagOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListTagOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListTagOptionsOutput{}
	fnCacher := func(out *servicecatalog.ListTagOptionsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListTagOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListTagOptionsPagesWithContext(ctx context.Context, input *servicecatalog.ListTagOptionsInput, fn func(*servicecatalog.ListTagOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.ListTagOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.ListTagOptionsOutput{}
	fnCacher := func(out *servicecatalog.ListTagOptionsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.ListTagOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) ListTagOptionsWithContext(ctx context.Context, input *servicecatalog.ListTagOptionsInput, opts ...request.Option) (*servicecatalog.ListTagOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ListTagOptionsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ListTagOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ScanProvisionedProducts(input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ScanProvisionedProductsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ScanProvisionedProducts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) ScanProvisionedProductsWithContext(ctx context.Context, input *servicecatalog.ScanProvisionedProductsInput, opts ...request.Option) (*servicecatalog.ScanProvisionedProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.ScanProvisionedProductsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.ScanProvisionedProductsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) SearchProducts(input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.SearchProductsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.SearchProducts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) SearchProductsAsAdmin(input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.SearchProductsAsAdminOutput), nil
	}
	out, err := c.ServiceCatalogAPI.SearchProductsAsAdmin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) SearchProductsAsAdminPages(input *servicecatalog.SearchProductsAsAdminInput, fn func(*servicecatalog.SearchProductsAsAdminOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.SearchProductsAsAdminOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.SearchProductsAsAdminOutput{}
	fnCacher := func(out *servicecatalog.SearchProductsAsAdminOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.SearchProductsAsAdminPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) SearchProductsAsAdminPagesWithContext(ctx context.Context, input *servicecatalog.SearchProductsAsAdminInput, fn func(*servicecatalog.SearchProductsAsAdminOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.SearchProductsAsAdminOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.SearchProductsAsAdminOutput{}
	fnCacher := func(out *servicecatalog.SearchProductsAsAdminOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.SearchProductsAsAdminPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) SearchProductsAsAdminWithContext(ctx context.Context, input *servicecatalog.SearchProductsAsAdminInput, opts ...request.Option) (*servicecatalog.SearchProductsAsAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.SearchProductsAsAdminOutput), nil
	}
	out, err := c.ServiceCatalogAPI.SearchProductsAsAdminWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) SearchProductsPages(input *servicecatalog.SearchProductsInput, fn func(*servicecatalog.SearchProductsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.SearchProductsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.SearchProductsOutput{}
	fnCacher := func(out *servicecatalog.SearchProductsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.SearchProductsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) SearchProductsPagesWithContext(ctx context.Context, input *servicecatalog.SearchProductsInput, fn func(*servicecatalog.SearchProductsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.SearchProductsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.SearchProductsOutput{}
	fnCacher := func(out *servicecatalog.SearchProductsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.SearchProductsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) SearchProductsWithContext(ctx context.Context, input *servicecatalog.SearchProductsInput, opts ...request.Option) (*servicecatalog.SearchProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.SearchProductsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.SearchProductsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) SearchProvisionedProducts(input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.SearchProvisionedProductsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.SearchProvisionedProducts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ServiceCatalog) SearchProvisionedProductsPages(input *servicecatalog.SearchProvisionedProductsInput, fn func(*servicecatalog.SearchProvisionedProductsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.SearchProvisionedProductsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.SearchProvisionedProductsOutput{}
	fnCacher := func(out *servicecatalog.SearchProvisionedProductsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.SearchProvisionedProductsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) SearchProvisionedProductsPagesWithContext(ctx context.Context, input *servicecatalog.SearchProvisionedProductsInput, fn func(*servicecatalog.SearchProvisionedProductsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*servicecatalog.SearchProvisionedProductsOutput), false)
		return nil
	}
	cachable := true
	output := &servicecatalog.SearchProvisionedProductsOutput{}
	fnCacher := func(out *servicecatalog.SearchProvisionedProductsOutput, more bool) bool {
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
	if err := c.ServiceCatalogAPI.SearchProvisionedProductsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ServiceCatalog) SearchProvisionedProductsWithContext(ctx context.Context, input *servicecatalog.SearchProvisionedProductsInput, opts ...request.Option) (*servicecatalog.SearchProvisionedProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*servicecatalog.SearchProvisionedProductsOutput), nil
	}
	out, err := c.ServiceCatalogAPI.SearchProvisionedProductsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

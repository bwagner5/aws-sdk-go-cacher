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
package textractcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/textract/textractiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Textract struct {
	textractiface.TextractAPI
	cache *cache.Cache
}

func New(textractapi textractiface.TextractAPI) *Textract {
	return &Textract{
		TextractAPI: textractapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Textract) GetDocumentAnalysis(input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetDocumentAnalysisOutput), nil
	}
	out, err := c.TextractAPI.GetDocumentAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetDocumentAnalysisWithContext(ctx context.Context, input *textract.GetDocumentAnalysisInput, opts ...request.Option) (*textract.GetDocumentAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetDocumentAnalysisOutput), nil
	}
	out, err := c.TextractAPI.GetDocumentAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetDocumentTextDetection(input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetDocumentTextDetectionOutput), nil
	}
	out, err := c.TextractAPI.GetDocumentTextDetection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetDocumentTextDetectionWithContext(ctx context.Context, input *textract.GetDocumentTextDetectionInput, opts ...request.Option) (*textract.GetDocumentTextDetectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetDocumentTextDetectionOutput), nil
	}
	out, err := c.TextractAPI.GetDocumentTextDetectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetExpenseAnalysis(input *textract.GetExpenseAnalysisInput) (*textract.GetExpenseAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetExpenseAnalysisOutput), nil
	}
	out, err := c.TextractAPI.GetExpenseAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetExpenseAnalysisWithContext(ctx context.Context, input *textract.GetExpenseAnalysisInput, opts ...request.Option) (*textract.GetExpenseAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetExpenseAnalysisOutput), nil
	}
	out, err := c.TextractAPI.GetExpenseAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetLendingAnalysis(input *textract.GetLendingAnalysisInput) (*textract.GetLendingAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetLendingAnalysisOutput), nil
	}
	out, err := c.TextractAPI.GetLendingAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetLendingAnalysisSummary(input *textract.GetLendingAnalysisSummaryInput) (*textract.GetLendingAnalysisSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetLendingAnalysisSummaryOutput), nil
	}
	out, err := c.TextractAPI.GetLendingAnalysisSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetLendingAnalysisSummaryWithContext(ctx context.Context, input *textract.GetLendingAnalysisSummaryInput, opts ...request.Option) (*textract.GetLendingAnalysisSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetLendingAnalysisSummaryOutput), nil
	}
	out, err := c.TextractAPI.GetLendingAnalysisSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Textract) GetLendingAnalysisWithContext(ctx context.Context, input *textract.GetLendingAnalysisInput, opts ...request.Option) (*textract.GetLendingAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*textract.GetLendingAnalysisOutput), nil
	}
	out, err := c.TextractAPI.GetLendingAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

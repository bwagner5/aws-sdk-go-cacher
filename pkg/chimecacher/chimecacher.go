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
package chimecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/chime"
	"github.com/aws/aws-sdk-go/service/chime/chimeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Chime struct {
	chimeiface.ChimeAPI
	cache *cache.Cache
}

func New(chimeapi chimeiface.ChimeAPI) *Chime {
	return &Chime{
		ChimeAPI: chimeapi,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Chime) BatchCreateAttendee(input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchCreateAttendeeOutput), nil
	}
	out, err := c.ChimeAPI.BatchCreateAttendee(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchCreateAttendeeWithContext(ctx context.Context, input *chime.BatchCreateAttendeeInput, opts ...request.Option) (*chime.BatchCreateAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchCreateAttendeeOutput), nil
	}
	out, err := c.ChimeAPI.BatchCreateAttendeeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchCreateChannelMembership(input *chime.BatchCreateChannelMembershipInput) (*chime.BatchCreateChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchCreateChannelMembershipOutput), nil
	}
	out, err := c.ChimeAPI.BatchCreateChannelMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchCreateChannelMembershipWithContext(ctx context.Context, input *chime.BatchCreateChannelMembershipInput, opts ...request.Option) (*chime.BatchCreateChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchCreateChannelMembershipOutput), nil
	}
	out, err := c.ChimeAPI.BatchCreateChannelMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchCreateRoomMembership(input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchCreateRoomMembershipOutput), nil
	}
	out, err := c.ChimeAPI.BatchCreateRoomMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchCreateRoomMembershipWithContext(ctx context.Context, input *chime.BatchCreateRoomMembershipInput, opts ...request.Option) (*chime.BatchCreateRoomMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchCreateRoomMembershipOutput), nil
	}
	out, err := c.ChimeAPI.BatchCreateRoomMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchDeletePhoneNumber(input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchDeletePhoneNumberOutput), nil
	}
	out, err := c.ChimeAPI.BatchDeletePhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchDeletePhoneNumberWithContext(ctx context.Context, input *chime.BatchDeletePhoneNumberInput, opts ...request.Option) (*chime.BatchDeletePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchDeletePhoneNumberOutput), nil
	}
	out, err := c.ChimeAPI.BatchDeletePhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchSuspendUser(input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchSuspendUserOutput), nil
	}
	out, err := c.ChimeAPI.BatchSuspendUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchSuspendUserWithContext(ctx context.Context, input *chime.BatchSuspendUserInput, opts ...request.Option) (*chime.BatchSuspendUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchSuspendUserOutput), nil
	}
	out, err := c.ChimeAPI.BatchSuspendUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchUnsuspendUser(input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchUnsuspendUserOutput), nil
	}
	out, err := c.ChimeAPI.BatchUnsuspendUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchUnsuspendUserWithContext(ctx context.Context, input *chime.BatchUnsuspendUserInput, opts ...request.Option) (*chime.BatchUnsuspendUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchUnsuspendUserOutput), nil
	}
	out, err := c.ChimeAPI.BatchUnsuspendUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchUpdatePhoneNumber(input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchUpdatePhoneNumberOutput), nil
	}
	out, err := c.ChimeAPI.BatchUpdatePhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchUpdatePhoneNumberWithContext(ctx context.Context, input *chime.BatchUpdatePhoneNumberInput, opts ...request.Option) (*chime.BatchUpdatePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchUpdatePhoneNumberOutput), nil
	}
	out, err := c.ChimeAPI.BatchUpdatePhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchUpdateUser(input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchUpdateUserOutput), nil
	}
	out, err := c.ChimeAPI.BatchUpdateUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) BatchUpdateUserWithContext(ctx context.Context, input *chime.BatchUpdateUserInput, opts ...request.Option) (*chime.BatchUpdateUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.BatchUpdateUserOutput), nil
	}
	out, err := c.ChimeAPI.BatchUpdateUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeAppInstance(input *chime.DescribeAppInstanceInput) (*chime.DescribeAppInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeAppInstanceOutput), nil
	}
	out, err := c.ChimeAPI.DescribeAppInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeAppInstanceAdmin(input *chime.DescribeAppInstanceAdminInput) (*chime.DescribeAppInstanceAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeAppInstanceAdminOutput), nil
	}
	out, err := c.ChimeAPI.DescribeAppInstanceAdmin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeAppInstanceAdminWithContext(ctx context.Context, input *chime.DescribeAppInstanceAdminInput, opts ...request.Option) (*chime.DescribeAppInstanceAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeAppInstanceAdminOutput), nil
	}
	out, err := c.ChimeAPI.DescribeAppInstanceAdminWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeAppInstanceUser(input *chime.DescribeAppInstanceUserInput) (*chime.DescribeAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.DescribeAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeAppInstanceUserWithContext(ctx context.Context, input *chime.DescribeAppInstanceUserInput, opts ...request.Option) (*chime.DescribeAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.DescribeAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeAppInstanceWithContext(ctx context.Context, input *chime.DescribeAppInstanceInput, opts ...request.Option) (*chime.DescribeAppInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeAppInstanceOutput), nil
	}
	out, err := c.ChimeAPI.DescribeAppInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannel(input *chime.DescribeChannelInput) (*chime.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelBan(input *chime.DescribeChannelBanInput) (*chime.DescribeChannelBanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelBanOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelBan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelBanWithContext(ctx context.Context, input *chime.DescribeChannelBanInput, opts ...request.Option) (*chime.DescribeChannelBanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelBanOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelBanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelMembership(input *chime.DescribeChannelMembershipInput) (*chime.DescribeChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelMembershipOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelMembershipForAppInstanceUser(input *chime.DescribeChannelMembershipForAppInstanceUserInput) (*chime.DescribeChannelMembershipForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelMembershipForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelMembershipForAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelMembershipForAppInstanceUserWithContext(ctx context.Context, input *chime.DescribeChannelMembershipForAppInstanceUserInput, opts ...request.Option) (*chime.DescribeChannelMembershipForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelMembershipForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelMembershipForAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelMembershipWithContext(ctx context.Context, input *chime.DescribeChannelMembershipInput, opts ...request.Option) (*chime.DescribeChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelMembershipOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelModeratedByAppInstanceUser(input *chime.DescribeChannelModeratedByAppInstanceUserInput) (*chime.DescribeChannelModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelModeratedByAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelModeratedByAppInstanceUserWithContext(ctx context.Context, input *chime.DescribeChannelModeratedByAppInstanceUserInput, opts ...request.Option) (*chime.DescribeChannelModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelModeratedByAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelModerator(input *chime.DescribeChannelModeratorInput) (*chime.DescribeChannelModeratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelModeratorOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelModerator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelModeratorWithContext(ctx context.Context, input *chime.DescribeChannelModeratorInput, opts ...request.Option) (*chime.DescribeChannelModeratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelModeratorOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelModeratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) DescribeChannelWithContext(ctx context.Context, input *chime.DescribeChannelInput, opts ...request.Option) (*chime.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.DescribeChannelOutput), nil
	}
	out, err := c.ChimeAPI.DescribeChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAccount(input *chime.GetAccountInput) (*chime.GetAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAccountOutput), nil
	}
	out, err := c.ChimeAPI.GetAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAccountSettings(input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAccountSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetAccountSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAccountSettingsWithContext(ctx context.Context, input *chime.GetAccountSettingsInput, opts ...request.Option) (*chime.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAccountSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetAccountSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAccountWithContext(ctx context.Context, input *chime.GetAccountInput, opts ...request.Option) (*chime.GetAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAccountOutput), nil
	}
	out, err := c.ChimeAPI.GetAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAppInstanceRetentionSettings(input *chime.GetAppInstanceRetentionSettingsInput) (*chime.GetAppInstanceRetentionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAppInstanceRetentionSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetAppInstanceRetentionSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAppInstanceRetentionSettingsWithContext(ctx context.Context, input *chime.GetAppInstanceRetentionSettingsInput, opts ...request.Option) (*chime.GetAppInstanceRetentionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAppInstanceRetentionSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetAppInstanceRetentionSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAppInstanceStreamingConfigurations(input *chime.GetAppInstanceStreamingConfigurationsInput) (*chime.GetAppInstanceStreamingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAppInstanceStreamingConfigurationsOutput), nil
	}
	out, err := c.ChimeAPI.GetAppInstanceStreamingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAppInstanceStreamingConfigurationsWithContext(ctx context.Context, input *chime.GetAppInstanceStreamingConfigurationsInput, opts ...request.Option) (*chime.GetAppInstanceStreamingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAppInstanceStreamingConfigurationsOutput), nil
	}
	out, err := c.ChimeAPI.GetAppInstanceStreamingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAttendee(input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAttendeeOutput), nil
	}
	out, err := c.ChimeAPI.GetAttendee(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetAttendeeWithContext(ctx context.Context, input *chime.GetAttendeeInput, opts ...request.Option) (*chime.GetAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetAttendeeOutput), nil
	}
	out, err := c.ChimeAPI.GetAttendeeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetBot(input *chime.GetBotInput) (*chime.GetBotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetBotOutput), nil
	}
	out, err := c.ChimeAPI.GetBot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetBotWithContext(ctx context.Context, input *chime.GetBotInput, opts ...request.Option) (*chime.GetBotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetBotOutput), nil
	}
	out, err := c.ChimeAPI.GetBotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetChannelMessage(input *chime.GetChannelMessageInput) (*chime.GetChannelMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetChannelMessageOutput), nil
	}
	out, err := c.ChimeAPI.GetChannelMessage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetChannelMessageWithContext(ctx context.Context, input *chime.GetChannelMessageInput, opts ...request.Option) (*chime.GetChannelMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetChannelMessageOutput), nil
	}
	out, err := c.ChimeAPI.GetChannelMessageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetEventsConfiguration(input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetEventsConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetEventsConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetEventsConfigurationWithContext(ctx context.Context, input *chime.GetEventsConfigurationInput, opts ...request.Option) (*chime.GetEventsConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetEventsConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetEventsConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetGlobalSettings(input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetGlobalSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetGlobalSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetGlobalSettingsWithContext(ctx context.Context, input *chime.GetGlobalSettingsInput, opts ...request.Option) (*chime.GetGlobalSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetGlobalSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetGlobalSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetMediaCapturePipeline(input *chime.GetMediaCapturePipelineInput) (*chime.GetMediaCapturePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetMediaCapturePipelineOutput), nil
	}
	out, err := c.ChimeAPI.GetMediaCapturePipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetMediaCapturePipelineWithContext(ctx context.Context, input *chime.GetMediaCapturePipelineInput, opts ...request.Option) (*chime.GetMediaCapturePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetMediaCapturePipelineOutput), nil
	}
	out, err := c.ChimeAPI.GetMediaCapturePipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetMeeting(input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetMeetingOutput), nil
	}
	out, err := c.ChimeAPI.GetMeeting(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetMeetingWithContext(ctx context.Context, input *chime.GetMeetingInput, opts ...request.Option) (*chime.GetMeetingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetMeetingOutput), nil
	}
	out, err := c.ChimeAPI.GetMeetingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetMessagingSessionEndpoint(input *chime.GetMessagingSessionEndpointInput) (*chime.GetMessagingSessionEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetMessagingSessionEndpointOutput), nil
	}
	out, err := c.ChimeAPI.GetMessagingSessionEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetMessagingSessionEndpointWithContext(ctx context.Context, input *chime.GetMessagingSessionEndpointInput, opts ...request.Option) (*chime.GetMessagingSessionEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetMessagingSessionEndpointOutput), nil
	}
	out, err := c.ChimeAPI.GetMessagingSessionEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetPhoneNumber(input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetPhoneNumberOutput), nil
	}
	out, err := c.ChimeAPI.GetPhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetPhoneNumberOrder(input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetPhoneNumberOrderOutput), nil
	}
	out, err := c.ChimeAPI.GetPhoneNumberOrder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetPhoneNumberOrderWithContext(ctx context.Context, input *chime.GetPhoneNumberOrderInput, opts ...request.Option) (*chime.GetPhoneNumberOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetPhoneNumberOrderOutput), nil
	}
	out, err := c.ChimeAPI.GetPhoneNumberOrderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetPhoneNumberSettings(input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetPhoneNumberSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetPhoneNumberSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetPhoneNumberSettingsWithContext(ctx context.Context, input *chime.GetPhoneNumberSettingsInput, opts ...request.Option) (*chime.GetPhoneNumberSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetPhoneNumberSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetPhoneNumberSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetPhoneNumberWithContext(ctx context.Context, input *chime.GetPhoneNumberInput, opts ...request.Option) (*chime.GetPhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetPhoneNumberOutput), nil
	}
	out, err := c.ChimeAPI.GetPhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetProxySession(input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetProxySessionOutput), nil
	}
	out, err := c.ChimeAPI.GetProxySession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetProxySessionWithContext(ctx context.Context, input *chime.GetProxySessionInput, opts ...request.Option) (*chime.GetProxySessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetProxySessionOutput), nil
	}
	out, err := c.ChimeAPI.GetProxySessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetRetentionSettings(input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetRetentionSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetRetentionSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetRetentionSettingsWithContext(ctx context.Context, input *chime.GetRetentionSettingsInput, opts ...request.Option) (*chime.GetRetentionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetRetentionSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetRetentionSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetRoom(input *chime.GetRoomInput) (*chime.GetRoomOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetRoomOutput), nil
	}
	out, err := c.ChimeAPI.GetRoom(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetRoomWithContext(ctx context.Context, input *chime.GetRoomInput, opts ...request.Option) (*chime.GetRoomOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetRoomOutput), nil
	}
	out, err := c.ChimeAPI.GetRoomWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetSipMediaApplication(input *chime.GetSipMediaApplicationInput) (*chime.GetSipMediaApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetSipMediaApplicationOutput), nil
	}
	out, err := c.ChimeAPI.GetSipMediaApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetSipMediaApplicationLoggingConfiguration(input *chime.GetSipMediaApplicationLoggingConfigurationInput) (*chime.GetSipMediaApplicationLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetSipMediaApplicationLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetSipMediaApplicationLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetSipMediaApplicationLoggingConfigurationWithContext(ctx context.Context, input *chime.GetSipMediaApplicationLoggingConfigurationInput, opts ...request.Option) (*chime.GetSipMediaApplicationLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetSipMediaApplicationLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetSipMediaApplicationLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetSipMediaApplicationWithContext(ctx context.Context, input *chime.GetSipMediaApplicationInput, opts ...request.Option) (*chime.GetSipMediaApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetSipMediaApplicationOutput), nil
	}
	out, err := c.ChimeAPI.GetSipMediaApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetSipRule(input *chime.GetSipRuleInput) (*chime.GetSipRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetSipRuleOutput), nil
	}
	out, err := c.ChimeAPI.GetSipRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetSipRuleWithContext(ctx context.Context, input *chime.GetSipRuleInput, opts ...request.Option) (*chime.GetSipRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetSipRuleOutput), nil
	}
	out, err := c.ChimeAPI.GetSipRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetUser(input *chime.GetUserInput) (*chime.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetUserOutput), nil
	}
	out, err := c.ChimeAPI.GetUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetUserSettings(input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetUserSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetUserSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetUserSettingsWithContext(ctx context.Context, input *chime.GetUserSettingsInput, opts ...request.Option) (*chime.GetUserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetUserSettingsOutput), nil
	}
	out, err := c.ChimeAPI.GetUserSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetUserWithContext(ctx context.Context, input *chime.GetUserInput, opts ...request.Option) (*chime.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetUserOutput), nil
	}
	out, err := c.ChimeAPI.GetUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnector(input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorEmergencyCallingConfiguration(input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorEmergencyCallingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorEmergencyCallingConfigurationWithContext(ctx context.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput, opts ...request.Option) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorEmergencyCallingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorGroup(input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorGroupOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorGroupWithContext(ctx context.Context, input *chime.GetVoiceConnectorGroupInput, opts ...request.Option) (*chime.GetVoiceConnectorGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorGroupOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorLoggingConfiguration(input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorLoggingConfigurationWithContext(ctx context.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput, opts ...request.Option) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorOrigination(input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorOriginationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorOrigination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorOriginationWithContext(ctx context.Context, input *chime.GetVoiceConnectorOriginationInput, opts ...request.Option) (*chime.GetVoiceConnectorOriginationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorOriginationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorOriginationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorProxy(input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorProxyOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorProxy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorProxyWithContext(ctx context.Context, input *chime.GetVoiceConnectorProxyInput, opts ...request.Option) (*chime.GetVoiceConnectorProxyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorProxyOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorProxyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorStreamingConfiguration(input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorStreamingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorStreamingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorStreamingConfigurationWithContext(ctx context.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput, opts ...request.Option) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorStreamingConfigurationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorStreamingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorTermination(input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorTerminationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorTermination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorTerminationHealth(input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorTerminationHealthOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorTerminationHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorTerminationHealthWithContext(ctx context.Context, input *chime.GetVoiceConnectorTerminationHealthInput, opts ...request.Option) (*chime.GetVoiceConnectorTerminationHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorTerminationHealthOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorTerminationHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorTerminationWithContext(ctx context.Context, input *chime.GetVoiceConnectorTerminationInput, opts ...request.Option) (*chime.GetVoiceConnectorTerminationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorTerminationOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorTerminationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) GetVoiceConnectorWithContext(ctx context.Context, input *chime.GetVoiceConnectorInput, opts ...request.Option) (*chime.GetVoiceConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.GetVoiceConnectorOutput), nil
	}
	out, err := c.ChimeAPI.GetVoiceConnectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAccounts(input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAccountsOutput), nil
	}
	out, err := c.ChimeAPI.ListAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAccountsPages(input *chime.ListAccountsInput, fn func(*chime.ListAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAccountsOutput{}
	fnCacher := func(out *chime.ListAccountsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAccountsPagesWithContext(ctx context.Context, input *chime.ListAccountsInput, fn func(*chime.ListAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAccountsOutput{}
	fnCacher := func(out *chime.ListAccountsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAccountsWithContext(ctx context.Context, input *chime.ListAccountsInput, opts ...request.Option) (*chime.ListAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAccountsOutput), nil
	}
	out, err := c.ChimeAPI.ListAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAppInstanceAdmins(input *chime.ListAppInstanceAdminsInput) (*chime.ListAppInstanceAdminsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAppInstanceAdminsOutput), nil
	}
	out, err := c.ChimeAPI.ListAppInstanceAdmins(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAppInstanceAdminsPages(input *chime.ListAppInstanceAdminsInput, fn func(*chime.ListAppInstanceAdminsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAppInstanceAdminsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAppInstanceAdminsOutput{}
	fnCacher := func(out *chime.ListAppInstanceAdminsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAppInstanceAdminsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAppInstanceAdminsPagesWithContext(ctx context.Context, input *chime.ListAppInstanceAdminsInput, fn func(*chime.ListAppInstanceAdminsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAppInstanceAdminsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAppInstanceAdminsOutput{}
	fnCacher := func(out *chime.ListAppInstanceAdminsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAppInstanceAdminsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAppInstanceAdminsWithContext(ctx context.Context, input *chime.ListAppInstanceAdminsInput, opts ...request.Option) (*chime.ListAppInstanceAdminsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAppInstanceAdminsOutput), nil
	}
	out, err := c.ChimeAPI.ListAppInstanceAdminsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAppInstanceUsers(input *chime.ListAppInstanceUsersInput) (*chime.ListAppInstanceUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAppInstanceUsersOutput), nil
	}
	out, err := c.ChimeAPI.ListAppInstanceUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAppInstanceUsersPages(input *chime.ListAppInstanceUsersInput, fn func(*chime.ListAppInstanceUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAppInstanceUsersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAppInstanceUsersOutput{}
	fnCacher := func(out *chime.ListAppInstanceUsersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAppInstanceUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAppInstanceUsersPagesWithContext(ctx context.Context, input *chime.ListAppInstanceUsersInput, fn func(*chime.ListAppInstanceUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAppInstanceUsersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAppInstanceUsersOutput{}
	fnCacher := func(out *chime.ListAppInstanceUsersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAppInstanceUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAppInstanceUsersWithContext(ctx context.Context, input *chime.ListAppInstanceUsersInput, opts ...request.Option) (*chime.ListAppInstanceUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAppInstanceUsersOutput), nil
	}
	out, err := c.ChimeAPI.ListAppInstanceUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAppInstances(input *chime.ListAppInstancesInput) (*chime.ListAppInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAppInstancesOutput), nil
	}
	out, err := c.ChimeAPI.ListAppInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAppInstancesPages(input *chime.ListAppInstancesInput, fn func(*chime.ListAppInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAppInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAppInstancesOutput{}
	fnCacher := func(out *chime.ListAppInstancesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAppInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAppInstancesPagesWithContext(ctx context.Context, input *chime.ListAppInstancesInput, fn func(*chime.ListAppInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAppInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAppInstancesOutput{}
	fnCacher := func(out *chime.ListAppInstancesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAppInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAppInstancesWithContext(ctx context.Context, input *chime.ListAppInstancesInput, opts ...request.Option) (*chime.ListAppInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAppInstancesOutput), nil
	}
	out, err := c.ChimeAPI.ListAppInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAttendeeTags(input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAttendeeTagsOutput), nil
	}
	out, err := c.ChimeAPI.ListAttendeeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAttendeeTagsWithContext(ctx context.Context, input *chime.ListAttendeeTagsInput, opts ...request.Option) (*chime.ListAttendeeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAttendeeTagsOutput), nil
	}
	out, err := c.ChimeAPI.ListAttendeeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAttendees(input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAttendeesOutput), nil
	}
	out, err := c.ChimeAPI.ListAttendees(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListAttendeesPages(input *chime.ListAttendeesInput, fn func(*chime.ListAttendeesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAttendeesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAttendeesOutput{}
	fnCacher := func(out *chime.ListAttendeesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAttendeesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAttendeesPagesWithContext(ctx context.Context, input *chime.ListAttendeesInput, fn func(*chime.ListAttendeesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListAttendeesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListAttendeesOutput{}
	fnCacher := func(out *chime.ListAttendeesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListAttendeesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListAttendeesWithContext(ctx context.Context, input *chime.ListAttendeesInput, opts ...request.Option) (*chime.ListAttendeesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListAttendeesOutput), nil
	}
	out, err := c.ChimeAPI.ListAttendeesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListBots(input *chime.ListBotsInput) (*chime.ListBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListBotsOutput), nil
	}
	out, err := c.ChimeAPI.ListBots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListBotsPages(input *chime.ListBotsInput, fn func(*chime.ListBotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListBotsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListBotsOutput{}
	fnCacher := func(out *chime.ListBotsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListBotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListBotsPagesWithContext(ctx context.Context, input *chime.ListBotsInput, fn func(*chime.ListBotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListBotsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListBotsOutput{}
	fnCacher := func(out *chime.ListBotsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListBotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListBotsWithContext(ctx context.Context, input *chime.ListBotsInput, opts ...request.Option) (*chime.ListBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListBotsOutput), nil
	}
	out, err := c.ChimeAPI.ListBotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelBans(input *chime.ListChannelBansInput) (*chime.ListChannelBansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelBansOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelBans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelBansPages(input *chime.ListChannelBansInput, fn func(*chime.ListChannelBansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelBansOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelBansOutput{}
	fnCacher := func(out *chime.ListChannelBansOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelBansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelBansPagesWithContext(ctx context.Context, input *chime.ListChannelBansInput, fn func(*chime.ListChannelBansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelBansOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelBansOutput{}
	fnCacher := func(out *chime.ListChannelBansOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelBansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelBansWithContext(ctx context.Context, input *chime.ListChannelBansInput, opts ...request.Option) (*chime.ListChannelBansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelBansOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelBansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelMemberships(input *chime.ListChannelMembershipsInput) (*chime.ListChannelMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelMembershipsOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelMembershipsForAppInstanceUser(input *chime.ListChannelMembershipsForAppInstanceUserInput) (*chime.ListChannelMembershipsForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelMembershipsForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelMembershipsForAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelMembershipsForAppInstanceUserPages(input *chime.ListChannelMembershipsForAppInstanceUserInput, fn func(*chime.ListChannelMembershipsForAppInstanceUserOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelMembershipsForAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelMembershipsForAppInstanceUserOutput{}
	fnCacher := func(out *chime.ListChannelMembershipsForAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelMembershipsForAppInstanceUserPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelMembershipsForAppInstanceUserPagesWithContext(ctx context.Context, input *chime.ListChannelMembershipsForAppInstanceUserInput, fn func(*chime.ListChannelMembershipsForAppInstanceUserOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelMembershipsForAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelMembershipsForAppInstanceUserOutput{}
	fnCacher := func(out *chime.ListChannelMembershipsForAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelMembershipsForAppInstanceUserPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelMembershipsForAppInstanceUserWithContext(ctx context.Context, input *chime.ListChannelMembershipsForAppInstanceUserInput, opts ...request.Option) (*chime.ListChannelMembershipsForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelMembershipsForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelMembershipsForAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelMembershipsPages(input *chime.ListChannelMembershipsInput, fn func(*chime.ListChannelMembershipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelMembershipsOutput{}
	fnCacher := func(out *chime.ListChannelMembershipsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelMembershipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelMembershipsPagesWithContext(ctx context.Context, input *chime.ListChannelMembershipsInput, fn func(*chime.ListChannelMembershipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelMembershipsOutput{}
	fnCacher := func(out *chime.ListChannelMembershipsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelMembershipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelMembershipsWithContext(ctx context.Context, input *chime.ListChannelMembershipsInput, opts ...request.Option) (*chime.ListChannelMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelMembershipsOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelMessages(input *chime.ListChannelMessagesInput) (*chime.ListChannelMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelMessagesOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelMessages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelMessagesPages(input *chime.ListChannelMessagesInput, fn func(*chime.ListChannelMessagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelMessagesOutput{}
	fnCacher := func(out *chime.ListChannelMessagesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelMessagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelMessagesPagesWithContext(ctx context.Context, input *chime.ListChannelMessagesInput, fn func(*chime.ListChannelMessagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelMessagesOutput{}
	fnCacher := func(out *chime.ListChannelMessagesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelMessagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelMessagesWithContext(ctx context.Context, input *chime.ListChannelMessagesInput, opts ...request.Option) (*chime.ListChannelMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelMessagesOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelMessagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelModerators(input *chime.ListChannelModeratorsInput) (*chime.ListChannelModeratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelModeratorsOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelModerators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelModeratorsPages(input *chime.ListChannelModeratorsInput, fn func(*chime.ListChannelModeratorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelModeratorsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelModeratorsOutput{}
	fnCacher := func(out *chime.ListChannelModeratorsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelModeratorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelModeratorsPagesWithContext(ctx context.Context, input *chime.ListChannelModeratorsInput, fn func(*chime.ListChannelModeratorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelModeratorsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelModeratorsOutput{}
	fnCacher := func(out *chime.ListChannelModeratorsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelModeratorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelModeratorsWithContext(ctx context.Context, input *chime.ListChannelModeratorsInput, opts ...request.Option) (*chime.ListChannelModeratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelModeratorsOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelModeratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannels(input *chime.ListChannelsInput) (*chime.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelsOutput), nil
	}
	out, err := c.ChimeAPI.ListChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelsModeratedByAppInstanceUser(input *chime.ListChannelsModeratedByAppInstanceUserInput) (*chime.ListChannelsModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelsModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelsModeratedByAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelsModeratedByAppInstanceUserPages(input *chime.ListChannelsModeratedByAppInstanceUserInput, fn func(*chime.ListChannelsModeratedByAppInstanceUserOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelsModeratedByAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelsModeratedByAppInstanceUserOutput{}
	fnCacher := func(out *chime.ListChannelsModeratedByAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelsModeratedByAppInstanceUserPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelsModeratedByAppInstanceUserPagesWithContext(ctx context.Context, input *chime.ListChannelsModeratedByAppInstanceUserInput, fn func(*chime.ListChannelsModeratedByAppInstanceUserOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelsModeratedByAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelsModeratedByAppInstanceUserOutput{}
	fnCacher := func(out *chime.ListChannelsModeratedByAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelsModeratedByAppInstanceUserPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelsModeratedByAppInstanceUserWithContext(ctx context.Context, input *chime.ListChannelsModeratedByAppInstanceUserInput, opts ...request.Option) (*chime.ListChannelsModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelsModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelsModeratedByAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListChannelsPages(input *chime.ListChannelsInput, fn func(*chime.ListChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelsOutput{}
	fnCacher := func(out *chime.ListChannelsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelsPagesWithContext(ctx context.Context, input *chime.ListChannelsInput, fn func(*chime.ListChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListChannelsOutput{}
	fnCacher := func(out *chime.ListChannelsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListChannelsWithContext(ctx context.Context, input *chime.ListChannelsInput, opts ...request.Option) (*chime.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListChannelsOutput), nil
	}
	out, err := c.ChimeAPI.ListChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListMediaCapturePipelines(input *chime.ListMediaCapturePipelinesInput) (*chime.ListMediaCapturePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListMediaCapturePipelinesOutput), nil
	}
	out, err := c.ChimeAPI.ListMediaCapturePipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListMediaCapturePipelinesPages(input *chime.ListMediaCapturePipelinesInput, fn func(*chime.ListMediaCapturePipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListMediaCapturePipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListMediaCapturePipelinesOutput{}
	fnCacher := func(out *chime.ListMediaCapturePipelinesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListMediaCapturePipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListMediaCapturePipelinesPagesWithContext(ctx context.Context, input *chime.ListMediaCapturePipelinesInput, fn func(*chime.ListMediaCapturePipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListMediaCapturePipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListMediaCapturePipelinesOutput{}
	fnCacher := func(out *chime.ListMediaCapturePipelinesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListMediaCapturePipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListMediaCapturePipelinesWithContext(ctx context.Context, input *chime.ListMediaCapturePipelinesInput, opts ...request.Option) (*chime.ListMediaCapturePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListMediaCapturePipelinesOutput), nil
	}
	out, err := c.ChimeAPI.ListMediaCapturePipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListMeetingTags(input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListMeetingTagsOutput), nil
	}
	out, err := c.ChimeAPI.ListMeetingTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListMeetingTagsWithContext(ctx context.Context, input *chime.ListMeetingTagsInput, opts ...request.Option) (*chime.ListMeetingTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListMeetingTagsOutput), nil
	}
	out, err := c.ChimeAPI.ListMeetingTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListMeetings(input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListMeetingsOutput), nil
	}
	out, err := c.ChimeAPI.ListMeetings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListMeetingsPages(input *chime.ListMeetingsInput, fn func(*chime.ListMeetingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListMeetingsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListMeetingsOutput{}
	fnCacher := func(out *chime.ListMeetingsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListMeetingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListMeetingsPagesWithContext(ctx context.Context, input *chime.ListMeetingsInput, fn func(*chime.ListMeetingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListMeetingsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListMeetingsOutput{}
	fnCacher := func(out *chime.ListMeetingsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListMeetingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListMeetingsWithContext(ctx context.Context, input *chime.ListMeetingsInput, opts ...request.Option) (*chime.ListMeetingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListMeetingsOutput), nil
	}
	out, err := c.ChimeAPI.ListMeetingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListPhoneNumberOrders(input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListPhoneNumberOrdersOutput), nil
	}
	out, err := c.ChimeAPI.ListPhoneNumberOrders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListPhoneNumberOrdersPages(input *chime.ListPhoneNumberOrdersInput, fn func(*chime.ListPhoneNumberOrdersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListPhoneNumberOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListPhoneNumberOrdersOutput{}
	fnCacher := func(out *chime.ListPhoneNumberOrdersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListPhoneNumberOrdersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListPhoneNumberOrdersPagesWithContext(ctx context.Context, input *chime.ListPhoneNumberOrdersInput, fn func(*chime.ListPhoneNumberOrdersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListPhoneNumberOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListPhoneNumberOrdersOutput{}
	fnCacher := func(out *chime.ListPhoneNumberOrdersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListPhoneNumberOrdersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListPhoneNumberOrdersWithContext(ctx context.Context, input *chime.ListPhoneNumberOrdersInput, opts ...request.Option) (*chime.ListPhoneNumberOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListPhoneNumberOrdersOutput), nil
	}
	out, err := c.ChimeAPI.ListPhoneNumberOrdersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListPhoneNumbers(input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListPhoneNumbersOutput), nil
	}
	out, err := c.ChimeAPI.ListPhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListPhoneNumbersPages(input *chime.ListPhoneNumbersInput, fn func(*chime.ListPhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListPhoneNumbersOutput{}
	fnCacher := func(out *chime.ListPhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListPhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListPhoneNumbersPagesWithContext(ctx context.Context, input *chime.ListPhoneNumbersInput, fn func(*chime.ListPhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListPhoneNumbersOutput{}
	fnCacher := func(out *chime.ListPhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListPhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListPhoneNumbersWithContext(ctx context.Context, input *chime.ListPhoneNumbersInput, opts ...request.Option) (*chime.ListPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListPhoneNumbersOutput), nil
	}
	out, err := c.ChimeAPI.ListPhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListProxySessions(input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListProxySessionsOutput), nil
	}
	out, err := c.ChimeAPI.ListProxySessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListProxySessionsPages(input *chime.ListProxySessionsInput, fn func(*chime.ListProxySessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListProxySessionsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListProxySessionsOutput{}
	fnCacher := func(out *chime.ListProxySessionsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListProxySessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListProxySessionsPagesWithContext(ctx context.Context, input *chime.ListProxySessionsInput, fn func(*chime.ListProxySessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListProxySessionsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListProxySessionsOutput{}
	fnCacher := func(out *chime.ListProxySessionsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListProxySessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListProxySessionsWithContext(ctx context.Context, input *chime.ListProxySessionsInput, opts ...request.Option) (*chime.ListProxySessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListProxySessionsOutput), nil
	}
	out, err := c.ChimeAPI.ListProxySessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListRoomMemberships(input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListRoomMembershipsOutput), nil
	}
	out, err := c.ChimeAPI.ListRoomMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListRoomMembershipsPages(input *chime.ListRoomMembershipsInput, fn func(*chime.ListRoomMembershipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListRoomMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListRoomMembershipsOutput{}
	fnCacher := func(out *chime.ListRoomMembershipsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListRoomMembershipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListRoomMembershipsPagesWithContext(ctx context.Context, input *chime.ListRoomMembershipsInput, fn func(*chime.ListRoomMembershipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListRoomMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListRoomMembershipsOutput{}
	fnCacher := func(out *chime.ListRoomMembershipsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListRoomMembershipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListRoomMembershipsWithContext(ctx context.Context, input *chime.ListRoomMembershipsInput, opts ...request.Option) (*chime.ListRoomMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListRoomMembershipsOutput), nil
	}
	out, err := c.ChimeAPI.ListRoomMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListRooms(input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListRoomsOutput), nil
	}
	out, err := c.ChimeAPI.ListRooms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListRoomsPages(input *chime.ListRoomsInput, fn func(*chime.ListRoomsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListRoomsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListRoomsOutput{}
	fnCacher := func(out *chime.ListRoomsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListRoomsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListRoomsPagesWithContext(ctx context.Context, input *chime.ListRoomsInput, fn func(*chime.ListRoomsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListRoomsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListRoomsOutput{}
	fnCacher := func(out *chime.ListRoomsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListRoomsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListRoomsWithContext(ctx context.Context, input *chime.ListRoomsInput, opts ...request.Option) (*chime.ListRoomsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListRoomsOutput), nil
	}
	out, err := c.ChimeAPI.ListRoomsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListSipMediaApplications(input *chime.ListSipMediaApplicationsInput) (*chime.ListSipMediaApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListSipMediaApplicationsOutput), nil
	}
	out, err := c.ChimeAPI.ListSipMediaApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListSipMediaApplicationsPages(input *chime.ListSipMediaApplicationsInput, fn func(*chime.ListSipMediaApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListSipMediaApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListSipMediaApplicationsOutput{}
	fnCacher := func(out *chime.ListSipMediaApplicationsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListSipMediaApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListSipMediaApplicationsPagesWithContext(ctx context.Context, input *chime.ListSipMediaApplicationsInput, fn func(*chime.ListSipMediaApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListSipMediaApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListSipMediaApplicationsOutput{}
	fnCacher := func(out *chime.ListSipMediaApplicationsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListSipMediaApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListSipMediaApplicationsWithContext(ctx context.Context, input *chime.ListSipMediaApplicationsInput, opts ...request.Option) (*chime.ListSipMediaApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListSipMediaApplicationsOutput), nil
	}
	out, err := c.ChimeAPI.ListSipMediaApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListSipRules(input *chime.ListSipRulesInput) (*chime.ListSipRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListSipRulesOutput), nil
	}
	out, err := c.ChimeAPI.ListSipRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListSipRulesPages(input *chime.ListSipRulesInput, fn func(*chime.ListSipRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListSipRulesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListSipRulesOutput{}
	fnCacher := func(out *chime.ListSipRulesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListSipRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListSipRulesPagesWithContext(ctx context.Context, input *chime.ListSipRulesInput, fn func(*chime.ListSipRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListSipRulesOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListSipRulesOutput{}
	fnCacher := func(out *chime.ListSipRulesOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListSipRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListSipRulesWithContext(ctx context.Context, input *chime.ListSipRulesInput, opts ...request.Option) (*chime.ListSipRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListSipRulesOutput), nil
	}
	out, err := c.ChimeAPI.ListSipRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListSupportedPhoneNumberCountries(input *chime.ListSupportedPhoneNumberCountriesInput) (*chime.ListSupportedPhoneNumberCountriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListSupportedPhoneNumberCountriesOutput), nil
	}
	out, err := c.ChimeAPI.ListSupportedPhoneNumberCountries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListSupportedPhoneNumberCountriesWithContext(ctx context.Context, input *chime.ListSupportedPhoneNumberCountriesInput, opts ...request.Option) (*chime.ListSupportedPhoneNumberCountriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListSupportedPhoneNumberCountriesOutput), nil
	}
	out, err := c.ChimeAPI.ListSupportedPhoneNumberCountriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListTagsForResource(input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListTagsForResourceWithContext(ctx context.Context, input *chime.ListTagsForResourceInput, opts ...request.Option) (*chime.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListUsers(input *chime.ListUsersInput) (*chime.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListUsersOutput), nil
	}
	out, err := c.ChimeAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListUsersPages(input *chime.ListUsersInput, fn func(*chime.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListUsersOutput{}
	fnCacher := func(out *chime.ListUsersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListUsersPagesWithContext(ctx context.Context, input *chime.ListUsersInput, fn func(*chime.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListUsersOutput{}
	fnCacher := func(out *chime.ListUsersOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListUsersWithContext(ctx context.Context, input *chime.ListUsersInput, opts ...request.Option) (*chime.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListUsersOutput), nil
	}
	out, err := c.ChimeAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListVoiceConnectorGroups(input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListVoiceConnectorGroupsOutput), nil
	}
	out, err := c.ChimeAPI.ListVoiceConnectorGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListVoiceConnectorGroupsPages(input *chime.ListVoiceConnectorGroupsInput, fn func(*chime.ListVoiceConnectorGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListVoiceConnectorGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListVoiceConnectorGroupsOutput{}
	fnCacher := func(out *chime.ListVoiceConnectorGroupsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListVoiceConnectorGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListVoiceConnectorGroupsPagesWithContext(ctx context.Context, input *chime.ListVoiceConnectorGroupsInput, fn func(*chime.ListVoiceConnectorGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListVoiceConnectorGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListVoiceConnectorGroupsOutput{}
	fnCacher := func(out *chime.ListVoiceConnectorGroupsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListVoiceConnectorGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListVoiceConnectorGroupsWithContext(ctx context.Context, input *chime.ListVoiceConnectorGroupsInput, opts ...request.Option) (*chime.ListVoiceConnectorGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListVoiceConnectorGroupsOutput), nil
	}
	out, err := c.ChimeAPI.ListVoiceConnectorGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListVoiceConnectorTerminationCredentials(input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListVoiceConnectorTerminationCredentialsOutput), nil
	}
	out, err := c.ChimeAPI.ListVoiceConnectorTerminationCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListVoiceConnectorTerminationCredentialsWithContext(ctx context.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput, opts ...request.Option) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListVoiceConnectorTerminationCredentialsOutput), nil
	}
	out, err := c.ChimeAPI.ListVoiceConnectorTerminationCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListVoiceConnectors(input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListVoiceConnectorsOutput), nil
	}
	out, err := c.ChimeAPI.ListVoiceConnectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) ListVoiceConnectorsPages(input *chime.ListVoiceConnectorsInput, fn func(*chime.ListVoiceConnectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListVoiceConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListVoiceConnectorsOutput{}
	fnCacher := func(out *chime.ListVoiceConnectorsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListVoiceConnectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListVoiceConnectorsPagesWithContext(ctx context.Context, input *chime.ListVoiceConnectorsInput, fn func(*chime.ListVoiceConnectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.ListVoiceConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &chime.ListVoiceConnectorsOutput{}
	fnCacher := func(out *chime.ListVoiceConnectorsOutput, more bool) bool {
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
	if err := c.ChimeAPI.ListVoiceConnectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) ListVoiceConnectorsWithContext(ctx context.Context, input *chime.ListVoiceConnectorsInput, opts ...request.Option) (*chime.ListVoiceConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.ListVoiceConnectorsOutput), nil
	}
	out, err := c.ChimeAPI.ListVoiceConnectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) SearchAvailablePhoneNumbers(input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.SearchAvailablePhoneNumbersOutput), nil
	}
	out, err := c.ChimeAPI.SearchAvailablePhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Chime) SearchAvailablePhoneNumbersPages(input *chime.SearchAvailablePhoneNumbersInput, fn func(*chime.SearchAvailablePhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.SearchAvailablePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.SearchAvailablePhoneNumbersOutput{}
	fnCacher := func(out *chime.SearchAvailablePhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeAPI.SearchAvailablePhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) SearchAvailablePhoneNumbersPagesWithContext(ctx context.Context, input *chime.SearchAvailablePhoneNumbersInput, fn func(*chime.SearchAvailablePhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chime.SearchAvailablePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chime.SearchAvailablePhoneNumbersOutput{}
	fnCacher := func(out *chime.SearchAvailablePhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeAPI.SearchAvailablePhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Chime) SearchAvailablePhoneNumbersWithContext(ctx context.Context, input *chime.SearchAvailablePhoneNumbersInput, opts ...request.Option) (*chime.SearchAvailablePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chime.SearchAvailablePhoneNumbersOutput), nil
	}
	out, err := c.ChimeAPI.SearchAvailablePhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}

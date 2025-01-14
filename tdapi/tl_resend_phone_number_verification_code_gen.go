// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ResendPhoneNumberVerificationCodeRequest represents TL type `resendPhoneNumberVerificationCode#51845ffc`.
type ResendPhoneNumberVerificationCodeRequest struct {
}

// ResendPhoneNumberVerificationCodeRequestTypeID is TL type id of ResendPhoneNumberVerificationCodeRequest.
const ResendPhoneNumberVerificationCodeRequestTypeID = 0x51845ffc

// Ensuring interfaces in compile-time for ResendPhoneNumberVerificationCodeRequest.
var (
	_ bin.Encoder     = &ResendPhoneNumberVerificationCodeRequest{}
	_ bin.Decoder     = &ResendPhoneNumberVerificationCodeRequest{}
	_ bin.BareEncoder = &ResendPhoneNumberVerificationCodeRequest{}
	_ bin.BareDecoder = &ResendPhoneNumberVerificationCodeRequest{}
)

func (r *ResendPhoneNumberVerificationCodeRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResendPhoneNumberVerificationCodeRequest) String() string {
	if r == nil {
		return "ResendPhoneNumberVerificationCodeRequest(nil)"
	}
	type Alias ResendPhoneNumberVerificationCodeRequest
	return fmt.Sprintf("ResendPhoneNumberVerificationCodeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResendPhoneNumberVerificationCodeRequest) TypeID() uint32 {
	return ResendPhoneNumberVerificationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResendPhoneNumberVerificationCodeRequest) TypeName() string {
	return "resendPhoneNumberVerificationCode"
}

// TypeInfo returns info about TL type.
func (r *ResendPhoneNumberVerificationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resendPhoneNumberVerificationCode",
		ID:   ResendPhoneNumberVerificationCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResendPhoneNumberVerificationCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendPhoneNumberVerificationCode#51845ffc as nil")
	}
	b.PutID(ResendPhoneNumberVerificationCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResendPhoneNumberVerificationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendPhoneNumberVerificationCode#51845ffc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResendPhoneNumberVerificationCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendPhoneNumberVerificationCode#51845ffc to nil")
	}
	if err := b.ConsumeID(ResendPhoneNumberVerificationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resendPhoneNumberVerificationCode#51845ffc: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResendPhoneNumberVerificationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendPhoneNumberVerificationCode#51845ffc to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResendPhoneNumberVerificationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resendPhoneNumberVerificationCode#51845ffc as nil")
	}
	b.ObjStart()
	b.PutID("resendPhoneNumberVerificationCode")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResendPhoneNumberVerificationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resendPhoneNumberVerificationCode#51845ffc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resendPhoneNumberVerificationCode"); err != nil {
				return fmt.Errorf("unable to decode resendPhoneNumberVerificationCode#51845ffc: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ResendPhoneNumberVerificationCode invokes method resendPhoneNumberVerificationCode#51845ffc returning error if any.
func (c *Client) ResendPhoneNumberVerificationCode(ctx context.Context) (*AuthenticationCodeInfo, error) {
	var result AuthenticationCodeInfo

	request := &ResendPhoneNumberVerificationCodeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// AccountGetSecureValueRequest represents TL type `account.getSecureValue#73665bc2`.
// Get saved Telegram Passport¹ document, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.getSecureValue for reference.
type AccountGetSecureValueRequest struct {
	// Requested value types
	Types []SecureValueTypeClass
}

// AccountGetSecureValueRequestTypeID is TL type id of AccountGetSecureValueRequest.
const AccountGetSecureValueRequestTypeID = 0x73665bc2

// Ensuring interfaces in compile-time for AccountGetSecureValueRequest.
var (
	_ bin.Encoder     = &AccountGetSecureValueRequest{}
	_ bin.Decoder     = &AccountGetSecureValueRequest{}
	_ bin.BareEncoder = &AccountGetSecureValueRequest{}
	_ bin.BareDecoder = &AccountGetSecureValueRequest{}
)

func (g *AccountGetSecureValueRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Types == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetSecureValueRequest) String() string {
	if g == nil {
		return "AccountGetSecureValueRequest(nil)"
	}
	type Alias AccountGetSecureValueRequest
	return fmt.Sprintf("AccountGetSecureValueRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetSecureValueRequest from given interface.
func (g *AccountGetSecureValueRequest) FillFrom(from interface {
	GetTypes() (value []SecureValueTypeClass)
}) {
	g.Types = from.GetTypes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetSecureValueRequest) TypeID() uint32 {
	return AccountGetSecureValueRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetSecureValueRequest) TypeName() string {
	return "account.getSecureValue"
}

// TypeInfo returns info about TL type.
func (g *AccountGetSecureValueRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getSecureValue",
		ID:   AccountGetSecureValueRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Types",
			SchemaName: "types",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetSecureValueRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getSecureValue#73665bc2 as nil")
	}
	b.PutID(AccountGetSecureValueRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetSecureValueRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getSecureValue#73665bc2 as nil")
	}
	b.PutVectorHeader(len(g.Types))
	for idx, v := range g.Types {
		if v == nil {
			return fmt.Errorf("unable to encode account.getSecureValue#73665bc2: field types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getSecureValue#73665bc2: field types element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetSecureValueRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getSecureValue#73665bc2 to nil")
	}
	if err := b.ConsumeID(AccountGetSecureValueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getSecureValue#73665bc2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetSecureValueRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getSecureValue#73665bc2 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.getSecureValue#73665bc2: field types: %w", err)
		}

		if headerLen > 0 {
			g.Types = make([]SecureValueTypeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureValueType(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.getSecureValue#73665bc2: field types: %w", err)
			}
			g.Types = append(g.Types, value)
		}
	}
	return nil
}

// GetTypes returns value of Types field.
func (g *AccountGetSecureValueRequest) GetTypes() (value []SecureValueTypeClass) {
	return g.Types
}

// MapTypes returns field Types wrapped in SecureValueTypeClassArray helper.
func (g *AccountGetSecureValueRequest) MapTypes() (value SecureValueTypeClassArray) {
	return SecureValueTypeClassArray(g.Types)
}

// AccountGetSecureValue invokes method account.getSecureValue#73665bc2 returning error if any.
// Get saved Telegram Passport¹ document, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.getSecureValue for reference.
func (c *Client) AccountGetSecureValue(ctx context.Context, types []SecureValueTypeClass) ([]SecureValue, error) {
	var result SecureValueVector

	request := &AccountGetSecureValueRequest{
		Types: types,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []SecureValue(result.Elems), nil
}

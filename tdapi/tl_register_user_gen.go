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

// RegisterUserRequest represents TL type `registerUser#f9719e1d`.
type RegisterUserRequest struct {
	// The first name of the user; 1-64 characters
	FirstName string
	// The last name of the user; 0-64 characters
	LastName string
}

// RegisterUserRequestTypeID is TL type id of RegisterUserRequest.
const RegisterUserRequestTypeID = 0xf9719e1d

// Ensuring interfaces in compile-time for RegisterUserRequest.
var (
	_ bin.Encoder     = &RegisterUserRequest{}
	_ bin.Decoder     = &RegisterUserRequest{}
	_ bin.BareEncoder = &RegisterUserRequest{}
	_ bin.BareDecoder = &RegisterUserRequest{}
)

func (r *RegisterUserRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.FirstName == "") {
		return false
	}
	if !(r.LastName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RegisterUserRequest) String() string {
	if r == nil {
		return "RegisterUserRequest(nil)"
	}
	type Alias RegisterUserRequest
	return fmt.Sprintf("RegisterUserRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RegisterUserRequest) TypeID() uint32 {
	return RegisterUserRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RegisterUserRequest) TypeName() string {
	return "registerUser"
}

// TypeInfo returns info about TL type.
func (r *RegisterUserRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "registerUser",
		ID:   RegisterUserRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RegisterUserRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode registerUser#f9719e1d as nil")
	}
	b.PutID(RegisterUserRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RegisterUserRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode registerUser#f9719e1d as nil")
	}
	b.PutString(r.FirstName)
	b.PutString(r.LastName)
	return nil
}

// Decode implements bin.Decoder.
func (r *RegisterUserRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode registerUser#f9719e1d to nil")
	}
	if err := b.ConsumeID(RegisterUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode registerUser#f9719e1d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RegisterUserRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode registerUser#f9719e1d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode registerUser#f9719e1d: field first_name: %w", err)
		}
		r.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode registerUser#f9719e1d: field last_name: %w", err)
		}
		r.LastName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RegisterUserRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode registerUser#f9719e1d as nil")
	}
	b.ObjStart()
	b.PutID("registerUser")
	b.FieldStart("first_name")
	b.PutString(r.FirstName)
	b.FieldStart("last_name")
	b.PutString(r.LastName)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RegisterUserRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode registerUser#f9719e1d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("registerUser"); err != nil {
				return fmt.Errorf("unable to decode registerUser#f9719e1d: %w", err)
			}
		case "first_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode registerUser#f9719e1d: field first_name: %w", err)
			}
			r.FirstName = value
		case "last_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode registerUser#f9719e1d: field last_name: %w", err)
			}
			r.LastName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFirstName returns value of FirstName field.
func (r *RegisterUserRequest) GetFirstName() (value string) {
	return r.FirstName
}

// GetLastName returns value of LastName field.
func (r *RegisterUserRequest) GetLastName() (value string) {
	return r.LastName
}

// RegisterUser invokes method registerUser#f9719e1d returning error if any.
func (c *Client) RegisterUser(ctx context.Context, request *RegisterUserRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

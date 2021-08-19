// Code generated by gotdgen, DO NOT EDIT.

package td

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

// SendMultipleSMSRequest represents TL type `sendMultipleSMS#df18e5ca`.
//
// See https://localhost:80/doc/constructor/sendMultipleSMS for reference.
type SendMultipleSMSRequest struct {
	// Messages field of SendMultipleSMSRequest.
	Messages []SMS
}

// SendMultipleSMSRequestTypeID is TL type id of SendMultipleSMSRequest.
const SendMultipleSMSRequestTypeID = 0xdf18e5ca

// Ensuring interfaces in compile-time for SendMultipleSMSRequest.
var (
	_ bin.Encoder     = &SendMultipleSMSRequest{}
	_ bin.Decoder     = &SendMultipleSMSRequest{}
	_ bin.BareEncoder = &SendMultipleSMSRequest{}
	_ bin.BareDecoder = &SendMultipleSMSRequest{}
)

func (s *SendMultipleSMSRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Messages == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMultipleSMSRequest) String() string {
	if s == nil {
		return "SendMultipleSMSRequest(nil)"
	}
	type Alias SendMultipleSMSRequest
	return fmt.Sprintf("SendMultipleSMSRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendMultipleSMSRequest) TypeID() uint32 {
	return SendMultipleSMSRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendMultipleSMSRequest) TypeName() string {
	return "sendMultipleSMS"
}

// TypeInfo returns info about TL type.
func (s *SendMultipleSMSRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendMultipleSMS",
		ID:   SendMultipleSMSRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendMultipleSMSRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMultipleSMS#df18e5ca as nil")
	}
	b.PutID(SendMultipleSMSRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendMultipleSMSRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMultipleSMS#df18e5ca as nil")
	}
	b.PutInt(len(s.Messages))
	for idx, v := range s.Messages {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sendMultipleSMS#df18e5ca: field messages element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMultipleSMSRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMultipleSMS#df18e5ca to nil")
	}
	if err := b.ConsumeID(SendMultipleSMSRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMultipleSMS#df18e5ca: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendMultipleSMSRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMultipleSMS#df18e5ca to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMultipleSMS#df18e5ca: field messages: %w", err)
		}

		if headerLen > 0 {
			s.Messages = make([]SMS, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SMS
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode sendMultipleSMS#df18e5ca: field messages: %w", err)
			}
			s.Messages = append(s.Messages, value)
		}
	}
	return nil
}

// GetMessages returns value of Messages field.
func (s *SendMultipleSMSRequest) GetMessages() (value []SMS) {
	return s.Messages
}

// SendMultipleSMS invokes method sendMultipleSMS#df18e5ca returning error if any.
//
// See https://localhost:80/doc/constructor/sendMultipleSMS for reference.
func (c *Client) SendMultipleSMS(ctx context.Context, messages []SMS) error {
	var ok Ok

	request := &SendMultipleSMSRequest{
		Messages: messages,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

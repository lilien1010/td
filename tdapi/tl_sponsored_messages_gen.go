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

// SponsoredMessages represents TL type `sponsoredMessages#4b9e4eb`.
type SponsoredMessages struct {
	// List of sponsored messages
	Messages []SponsoredMessage
}

// SponsoredMessagesTypeID is TL type id of SponsoredMessages.
const SponsoredMessagesTypeID = 0x4b9e4eb

// Ensuring interfaces in compile-time for SponsoredMessages.
var (
	_ bin.Encoder     = &SponsoredMessages{}
	_ bin.Decoder     = &SponsoredMessages{}
	_ bin.BareEncoder = &SponsoredMessages{}
	_ bin.BareDecoder = &SponsoredMessages{}
)

func (s *SponsoredMessages) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Messages == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredMessages) String() string {
	if s == nil {
		return "SponsoredMessages(nil)"
	}
	type Alias SponsoredMessages
	return fmt.Sprintf("SponsoredMessages%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredMessages) TypeID() uint32 {
	return SponsoredMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredMessages) TypeName() string {
	return "sponsoredMessages"
}

// TypeInfo returns info about TL type.
func (s *SponsoredMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredMessages",
		ID:   SponsoredMessagesTypeID,
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
func (s *SponsoredMessages) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessages#4b9e4eb as nil")
	}
	b.PutID(SponsoredMessagesTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredMessages) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessages#4b9e4eb as nil")
	}
	b.PutInt(len(s.Messages))
	for idx, v := range s.Messages {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare sponsoredMessages#4b9e4eb: field messages element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredMessages) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessages#4b9e4eb to nil")
	}
	if err := b.ConsumeID(SponsoredMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredMessages#4b9e4eb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredMessages) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessages#4b9e4eb to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessages#4b9e4eb: field messages: %w", err)
		}

		if headerLen > 0 {
			s.Messages = make([]SponsoredMessage, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SponsoredMessage
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare sponsoredMessages#4b9e4eb: field messages: %w", err)
			}
			s.Messages = append(s.Messages, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SponsoredMessages) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessages#4b9e4eb as nil")
	}
	b.ObjStart()
	b.PutID("sponsoredMessages")
	b.FieldStart("messages")
	b.ArrStart()
	for idx, v := range s.Messages {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredMessages#4b9e4eb: field messages element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SponsoredMessages) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessages#4b9e4eb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sponsoredMessages"); err != nil {
				return fmt.Errorf("unable to decode sponsoredMessages#4b9e4eb: %w", err)
			}
		case "messages":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value SponsoredMessage
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode sponsoredMessages#4b9e4eb: field messages: %w", err)
				}
				s.Messages = append(s.Messages, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode sponsoredMessages#4b9e4eb: field messages: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessages returns value of Messages field.
func (s *SponsoredMessages) GetMessages() (value []SponsoredMessage) {
	return s.Messages
}

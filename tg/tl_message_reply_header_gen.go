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

// MessageReplyHeader represents TL type `messageReplyHeader#a6d57763`.
// Message replies and thread¹ information
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/constructor/messageReplyHeader for reference.
type MessageReplyHeader struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// ID of message to which this message is replying
	ReplyToMsgID int
	// For replies sent in channel discussion threads¹ of which the current user is not a
	// member, the discussion group ID
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplyToPeerID and GetReplyToPeerID helpers.
	ReplyToPeerID PeerClass
	// ID of the message that started this message thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplyToTopID and GetReplyToTopID helpers.
	ReplyToTopID int
}

// MessageReplyHeaderTypeID is TL type id of MessageReplyHeader.
const MessageReplyHeaderTypeID = 0xa6d57763

// Ensuring interfaces in compile-time for MessageReplyHeader.
var (
	_ bin.Encoder     = &MessageReplyHeader{}
	_ bin.Decoder     = &MessageReplyHeader{}
	_ bin.BareEncoder = &MessageReplyHeader{}
	_ bin.BareDecoder = &MessageReplyHeader{}
)

func (m *MessageReplyHeader) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.ReplyToMsgID == 0) {
		return false
	}
	if !(m.ReplyToPeerID == nil) {
		return false
	}
	if !(m.ReplyToTopID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyHeader) String() string {
	if m == nil {
		return "MessageReplyHeader(nil)"
	}
	type Alias MessageReplyHeader
	return fmt.Sprintf("MessageReplyHeader%+v", Alias(*m))
}

// FillFrom fills MessageReplyHeader from given interface.
func (m *MessageReplyHeader) FillFrom(from interface {
	GetReplyToMsgID() (value int)
	GetReplyToPeerID() (value PeerClass, ok bool)
	GetReplyToTopID() (value int, ok bool)
}) {
	m.ReplyToMsgID = from.GetReplyToMsgID()
	if val, ok := from.GetReplyToPeerID(); ok {
		m.ReplyToPeerID = val
	}

	if val, ok := from.GetReplyToTopID(); ok {
		m.ReplyToTopID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplyHeader) TypeID() uint32 {
	return MessageReplyHeaderTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplyHeader) TypeName() string {
	return "messageReplyHeader"
}

// TypeInfo returns info about TL type.
func (m *MessageReplyHeader) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplyHeader",
		ID:   MessageReplyHeaderTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReplyToMsgID",
			SchemaName: "reply_to_msg_id",
		},
		{
			Name:       "ReplyToPeerID",
			SchemaName: "reply_to_peer_id",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "ReplyToTopID",
			SchemaName: "reply_to_top_id",
			Null:       !m.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageReplyHeader) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyHeader#a6d57763 as nil")
	}
	b.PutID(MessageReplyHeaderTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplyHeader) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyHeader#a6d57763 as nil")
	}
	if !(m.ReplyToPeerID == nil) {
		m.Flags.Set(0)
	}
	if !(m.ReplyToTopID == 0) {
		m.Flags.Set(1)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyHeader#a6d57763: field flags: %w", err)
	}
	b.PutInt(m.ReplyToMsgID)
	if m.Flags.Has(0) {
		if m.ReplyToPeerID == nil {
			return fmt.Errorf("unable to encode messageReplyHeader#a6d57763: field reply_to_peer_id is nil")
		}
		if err := m.ReplyToPeerID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageReplyHeader#a6d57763: field reply_to_peer_id: %w", err)
		}
	}
	if m.Flags.Has(1) {
		b.PutInt(m.ReplyToTopID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplyHeader) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyHeader#a6d57763 to nil")
	}
	if err := b.ConsumeID(MessageReplyHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplyHeader) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyHeader#a6d57763 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field reply_to_msg_id: %w", err)
		}
		m.ReplyToMsgID = value
	}
	if m.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field reply_to_peer_id: %w", err)
		}
		m.ReplyToPeerID = value
	}
	if m.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field reply_to_top_id: %w", err)
		}
		m.ReplyToTopID = value
	}
	return nil
}

// GetReplyToMsgID returns value of ReplyToMsgID field.
func (m *MessageReplyHeader) GetReplyToMsgID() (value int) {
	return m.ReplyToMsgID
}

// SetReplyToPeerID sets value of ReplyToPeerID conditional field.
func (m *MessageReplyHeader) SetReplyToPeerID(value PeerClass) {
	m.Flags.Set(0)
	m.ReplyToPeerID = value
}

// GetReplyToPeerID returns value of ReplyToPeerID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToPeerID() (value PeerClass, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.ReplyToPeerID, true
}

// SetReplyToTopID sets value of ReplyToTopID conditional field.
func (m *MessageReplyHeader) SetReplyToTopID(value int) {
	m.Flags.Set(1)
	m.ReplyToTopID = value
}

// GetReplyToTopID returns value of ReplyToTopID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToTopID() (value int, ok bool) {
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.ReplyToTopID, true
}

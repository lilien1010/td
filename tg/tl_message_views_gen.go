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

// MessageViews represents TL type `messageViews#455b853d`.
// View, forward counter + info about replies of a specific message
//
// See https://core.telegram.org/constructor/messageViews for reference.
type MessageViews struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Viewcount of message
	//
	// Use SetViews and GetViews helpers.
	Views int
	// Forward count of message
	//
	// Use SetForwards and GetForwards helpers.
	Forwards int
	// Reply and thread¹ information of message
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplies and GetReplies helpers.
	Replies MessageReplies
}

// MessageViewsTypeID is TL type id of MessageViews.
const MessageViewsTypeID = 0x455b853d

// Ensuring interfaces in compile-time for MessageViews.
var (
	_ bin.Encoder     = &MessageViews{}
	_ bin.Decoder     = &MessageViews{}
	_ bin.BareEncoder = &MessageViews{}
	_ bin.BareDecoder = &MessageViews{}
)

func (m *MessageViews) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Views == 0) {
		return false
	}
	if !(m.Forwards == 0) {
		return false
	}
	if !(m.Replies.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageViews) String() string {
	if m == nil {
		return "MessageViews(nil)"
	}
	type Alias MessageViews
	return fmt.Sprintf("MessageViews%+v", Alias(*m))
}

// FillFrom fills MessageViews from given interface.
func (m *MessageViews) FillFrom(from interface {
	GetViews() (value int, ok bool)
	GetForwards() (value int, ok bool)
	GetReplies() (value MessageReplies, ok bool)
}) {
	if val, ok := from.GetViews(); ok {
		m.Views = val
	}

	if val, ok := from.GetForwards(); ok {
		m.Forwards = val
	}

	if val, ok := from.GetReplies(); ok {
		m.Replies = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageViews) TypeID() uint32 {
	return MessageViewsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageViews) TypeName() string {
	return "messageViews"
}

// TypeInfo returns info about TL type.
func (m *MessageViews) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageViews",
		ID:   MessageViewsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Views",
			SchemaName: "views",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "Forwards",
			SchemaName: "forwards",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "Replies",
			SchemaName: "replies",
			Null:       !m.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageViews) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageViews#455b853d as nil")
	}
	b.PutID(MessageViewsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageViews) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageViews#455b853d as nil")
	}
	if !(m.Views == 0) {
		m.Flags.Set(0)
	}
	if !(m.Forwards == 0) {
		m.Flags.Set(1)
	}
	if !(m.Replies.Zero()) {
		m.Flags.Set(2)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageViews#455b853d: field flags: %w", err)
	}
	if m.Flags.Has(0) {
		b.PutInt(m.Views)
	}
	if m.Flags.Has(1) {
		b.PutInt(m.Forwards)
	}
	if m.Flags.Has(2) {
		if err := m.Replies.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageViews#455b853d: field replies: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageViews) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageViews#455b853d to nil")
	}
	if err := b.ConsumeID(MessageViewsTypeID); err != nil {
		return fmt.Errorf("unable to decode messageViews#455b853d: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageViews) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageViews#455b853d to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field flags: %w", err)
		}
	}
	if m.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field views: %w", err)
		}
		m.Views = value
	}
	if m.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field forwards: %w", err)
		}
		m.Forwards = value
	}
	if m.Flags.Has(2) {
		if err := m.Replies.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageViews#455b853d: field replies: %w", err)
		}
	}
	return nil
}

// SetViews sets value of Views conditional field.
func (m *MessageViews) SetViews(value int) {
	m.Flags.Set(0)
	m.Views = value
}

// GetViews returns value of Views conditional field and
// boolean which is true if field was set.
func (m *MessageViews) GetViews() (value int, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.Views, true
}

// SetForwards sets value of Forwards conditional field.
func (m *MessageViews) SetForwards(value int) {
	m.Flags.Set(1)
	m.Forwards = value
}

// GetForwards returns value of Forwards conditional field and
// boolean which is true if field was set.
func (m *MessageViews) GetForwards() (value int, ok bool) {
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.Forwards, true
}

// SetReplies sets value of Replies conditional field.
func (m *MessageViews) SetReplies(value MessageReplies) {
	m.Flags.Set(2)
	m.Replies = value
}

// GetReplies returns value of Replies conditional field and
// boolean which is true if field was set.
func (m *MessageViews) GetReplies() (value MessageReplies, ok bool) {
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.Replies, true
}

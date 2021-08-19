// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// MsgsAck represents TL type `msgs_ack#62d6b459`.
type MsgsAck struct {
	// MsgIDs field of MsgsAck.
	MsgIDs []int64
}

// MsgsAckTypeID is TL type id of MsgsAck.
const MsgsAckTypeID = 0x62d6b459

// Ensuring interfaces in compile-time for MsgsAck.
var (
	_ bin.Encoder     = &MsgsAck{}
	_ bin.Decoder     = &MsgsAck{}
	_ bin.BareEncoder = &MsgsAck{}
	_ bin.BareDecoder = &MsgsAck{}
)

func (m *MsgsAck) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsAck) String() string {
	if m == nil {
		return "MsgsAck(nil)"
	}
	type Alias MsgsAck
	return fmt.Sprintf("MsgsAck%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgsAck) TypeID() uint32 {
	return MsgsAckTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgsAck) TypeName() string {
	return "msgs_ack"
}

// TypeInfo returns info about TL type.
func (m *MsgsAck) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msgs_ack",
		ID:   MsgsAckTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgIDs",
			SchemaName: "msg_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgsAck) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_ack#62d6b459 as nil")
	}
	b.PutID(MsgsAckTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgsAck) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_ack#62d6b459 as nil")
	}
	b.PutVectorHeader(len(m.MsgIDs))
	for _, v := range m.MsgIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgsAck) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_ack#62d6b459 to nil")
	}
	if err := b.ConsumeID(MsgsAckTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_ack#62d6b459: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgsAck) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_ack#62d6b459 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_ack#62d6b459: field msg_ids: %w", err)
		}

		if headerLen > 0 {
			m.MsgIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msgs_ack#62d6b459: field msg_ids: %w", err)
			}
			m.MsgIDs = append(m.MsgIDs, value)
		}
	}
	return nil
}

// GetMsgIDs returns value of MsgIDs field.
func (m *MsgsAck) GetMsgIDs() (value []int64) {
	return m.MsgIDs
}

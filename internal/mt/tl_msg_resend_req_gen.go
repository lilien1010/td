// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// MsgResendReq represents TL type `msg_resend_req#7d861a08`.
type MsgResendReq struct {
	// MsgIDs field of MsgResendReq.
	MsgIDs []int64
}

// MsgResendReqTypeID is TL type id of MsgResendReq.
const MsgResendReqTypeID = 0x7d861a08

func (m *MsgResendReq) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgResendReq) String() string {
	if m == nil {
		return "MsgResendReq(nil)"
	}
	type Alias MsgResendReq
	return fmt.Sprintf("MsgResendReq%+v", Alias(*m))
}

// FillFrom fills MsgResendReq from given interface.
func (m *MsgResendReq) FillFrom(from interface {
	GetMsgIDs() (value []int64)
}) {
	m.MsgIDs = from.GetMsgIDs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgResendReq) TypeID() uint32 {
	return MsgResendReqTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgResendReq) TypeName() string {
	return "msg_resend_req"
}

// TypeInfo returns info about TL type.
func (m *MsgResendReq) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msg_resend_req",
		ID:   MsgResendReqTypeID,
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
func (m *MsgResendReq) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_resend_req#7d861a08 as nil")
	}
	b.PutID(MsgResendReqTypeID)
	b.PutVectorHeader(len(m.MsgIDs))
	for _, v := range m.MsgIDs {
		b.PutLong(v)
	}
	return nil
}

// GetMsgIDs returns value of MsgIDs field.
func (m *MsgResendReq) GetMsgIDs() (value []int64) {
	return m.MsgIDs
}

// Decode implements bin.Decoder.
func (m *MsgResendReq) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_resend_req#7d861a08 to nil")
	}
	if err := b.ConsumeID(MsgResendReqTypeID); err != nil {
		return fmt.Errorf("unable to decode msg_resend_req#7d861a08: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msg_resend_req#7d861a08: field msg_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msg_resend_req#7d861a08: field msg_ids: %w", err)
			}
			m.MsgIDs = append(m.MsgIDs, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MsgResendReq.
var (
	_ bin.Encoder = &MsgResendReq{}
	_ bin.Decoder = &MsgResendReq{}
)

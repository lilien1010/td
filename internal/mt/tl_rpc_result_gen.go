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

// RPCResult represents TL type `rpc_result#f35c6d01`.
type RPCResult struct {
	// ReqMsgID field of RPCResult.
	ReqMsgID int64
	// Result field of RPCResult.
	Result GzipPacked
}

// RPCResultTypeID is TL type id of RPCResult.
const RPCResultTypeID = 0xf35c6d01

// Ensuring interfaces in compile-time for RPCResult.
var (
	_ bin.Encoder     = &RPCResult{}
	_ bin.Decoder     = &RPCResult{}
	_ bin.BareEncoder = &RPCResult{}
	_ bin.BareDecoder = &RPCResult{}
)

func (r *RPCResult) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ReqMsgID == 0) {
		return false
	}
	if !(r.Result.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCResult) String() string {
	if r == nil {
		return "RPCResult(nil)"
	}
	type Alias RPCResult
	return fmt.Sprintf("RPCResult%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCResult) TypeID() uint32 {
	return RPCResultTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCResult) TypeName() string {
	return "rpc_result"
}

// TypeInfo returns info about TL type.
func (r *RPCResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_result",
		ID:   RPCResultTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReqMsgID",
			SchemaName: "req_msg_id",
		},
		{
			Name:       "Result",
			SchemaName: "result",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCResult) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_result#f35c6d01 as nil")
	}
	b.PutID(RPCResultTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RPCResult) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_result#f35c6d01 as nil")
	}
	b.PutLong(r.ReqMsgID)
	if err := r.Result.Encode(b); err != nil {
		return fmt.Errorf("unable to encode rpc_result#f35c6d01: field result: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCResult) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_result#f35c6d01 to nil")
	}
	if err := b.ConsumeID(RPCResultTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_result#f35c6d01: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RPCResult) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_result#f35c6d01 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_result#f35c6d01: field req_msg_id: %w", err)
		}
		r.ReqMsgID = value
	}
	{
		if err := r.Result.Decode(b); err != nil {
			return fmt.Errorf("unable to decode rpc_result#f35c6d01: field result: %w", err)
		}
	}
	return nil
}

// GetReqMsgID returns value of ReqMsgID field.
func (r *RPCResult) GetReqMsgID() (value int64) {
	return r.ReqMsgID
}

// GetResult returns value of Result field.
func (r *RPCResult) GetResult() (value GzipPacked) {
	return r.Result
}

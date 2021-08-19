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

// HelpGetNearestDCRequest represents TL type `help.getNearestDc#1fb33026`.
// Returns info on data centre nearest to the user.
//
// See https://core.telegram.org/method/help.getNearestDc for reference.
type HelpGetNearestDCRequest struct {
}

// HelpGetNearestDCRequestTypeID is TL type id of HelpGetNearestDCRequest.
const HelpGetNearestDCRequestTypeID = 0x1fb33026

// Ensuring interfaces in compile-time for HelpGetNearestDCRequest.
var (
	_ bin.Encoder     = &HelpGetNearestDCRequest{}
	_ bin.Decoder     = &HelpGetNearestDCRequest{}
	_ bin.BareEncoder = &HelpGetNearestDCRequest{}
	_ bin.BareDecoder = &HelpGetNearestDCRequest{}
)

func (g *HelpGetNearestDCRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetNearestDCRequest) String() string {
	if g == nil {
		return "HelpGetNearestDCRequest(nil)"
	}
	type Alias HelpGetNearestDCRequest
	return fmt.Sprintf("HelpGetNearestDCRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetNearestDCRequest) TypeID() uint32 {
	return HelpGetNearestDCRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetNearestDCRequest) TypeName() string {
	return "help.getNearestDc"
}

// TypeInfo returns info about TL type.
func (g *HelpGetNearestDCRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getNearestDc",
		ID:   HelpGetNearestDCRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetNearestDCRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getNearestDc#1fb33026 as nil")
	}
	b.PutID(HelpGetNearestDCRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetNearestDCRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getNearestDc#1fb33026 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetNearestDCRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getNearestDc#1fb33026 to nil")
	}
	if err := b.ConsumeID(HelpGetNearestDCRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getNearestDc#1fb33026: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetNearestDCRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getNearestDc#1fb33026 to nil")
	}
	return nil
}

// HelpGetNearestDC invokes method help.getNearestDc#1fb33026 returning error if any.
// Returns info on data centre nearest to the user.
//
// See https://core.telegram.org/method/help.getNearestDc for reference.
func (c *Client) HelpGetNearestDC(ctx context.Context) (*NearestDC, error) {
	var result NearestDC

	request := &HelpGetNearestDCRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

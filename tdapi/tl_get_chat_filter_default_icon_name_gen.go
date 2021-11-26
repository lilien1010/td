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

// GetChatFilterDefaultIconNameRequest represents TL type `getChatFilterDefaultIconName#b023d638`.
type GetChatFilterDefaultIconNameRequest struct {
	// Chat filter
	Filter ChatFilter
}

// GetChatFilterDefaultIconNameRequestTypeID is TL type id of GetChatFilterDefaultIconNameRequest.
const GetChatFilterDefaultIconNameRequestTypeID = 0xb023d638

// Ensuring interfaces in compile-time for GetChatFilterDefaultIconNameRequest.
var (
	_ bin.Encoder     = &GetChatFilterDefaultIconNameRequest{}
	_ bin.Decoder     = &GetChatFilterDefaultIconNameRequest{}
	_ bin.BareEncoder = &GetChatFilterDefaultIconNameRequest{}
	_ bin.BareDecoder = &GetChatFilterDefaultIconNameRequest{}
)

func (g *GetChatFilterDefaultIconNameRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Filter.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatFilterDefaultIconNameRequest) String() string {
	if g == nil {
		return "GetChatFilterDefaultIconNameRequest(nil)"
	}
	type Alias GetChatFilterDefaultIconNameRequest
	return fmt.Sprintf("GetChatFilterDefaultIconNameRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatFilterDefaultIconNameRequest) TypeID() uint32 {
	return GetChatFilterDefaultIconNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatFilterDefaultIconNameRequest) TypeName() string {
	return "getChatFilterDefaultIconName"
}

// TypeInfo returns info about TL type.
func (g *GetChatFilterDefaultIconNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatFilterDefaultIconName",
		ID:   GetChatFilterDefaultIconNameRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatFilterDefaultIconNameRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFilterDefaultIconName#b023d638 as nil")
	}
	b.PutID(GetChatFilterDefaultIconNameRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatFilterDefaultIconNameRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFilterDefaultIconName#b023d638 as nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatFilterDefaultIconName#b023d638: field filter: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatFilterDefaultIconNameRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFilterDefaultIconName#b023d638 to nil")
	}
	if err := b.ConsumeID(GetChatFilterDefaultIconNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatFilterDefaultIconName#b023d638: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatFilterDefaultIconNameRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFilterDefaultIconName#b023d638 to nil")
	}
	{
		if err := g.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getChatFilterDefaultIconName#b023d638: field filter: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatFilterDefaultIconNameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFilterDefaultIconName#b023d638 as nil")
	}
	b.ObjStart()
	b.PutID("getChatFilterDefaultIconName")
	b.FieldStart("filter")
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatFilterDefaultIconName#b023d638: field filter: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatFilterDefaultIconNameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFilterDefaultIconName#b023d638 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatFilterDefaultIconName"); err != nil {
				return fmt.Errorf("unable to decode getChatFilterDefaultIconName#b023d638: %w", err)
			}
		case "filter":
			if err := g.Filter.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getChatFilterDefaultIconName#b023d638: field filter: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFilter returns value of Filter field.
func (g *GetChatFilterDefaultIconNameRequest) GetFilter() (value ChatFilter) {
	return g.Filter
}

// GetChatFilterDefaultIconName invokes method getChatFilterDefaultIconName#b023d638 returning error if any.
func (c *Client) GetChatFilterDefaultIconName(ctx context.Context, filter ChatFilter) (*Text, error) {
	var result Text

	request := &GetChatFilterDefaultIconNameRequest{
		Filter: filter,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
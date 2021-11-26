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

// GetMessageEmbeddingCodeRequest represents TL type `getMessageEmbeddingCode#62a4cd09`.
type GetMessageEmbeddingCodeRequest struct {
	// Identifier of the chat to which the message belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// Pass true to return an HTML code for embedding of the whole media album
	ForAlbum bool
}

// GetMessageEmbeddingCodeRequestTypeID is TL type id of GetMessageEmbeddingCodeRequest.
const GetMessageEmbeddingCodeRequestTypeID = 0x62a4cd09

// Ensuring interfaces in compile-time for GetMessageEmbeddingCodeRequest.
var (
	_ bin.Encoder     = &GetMessageEmbeddingCodeRequest{}
	_ bin.Decoder     = &GetMessageEmbeddingCodeRequest{}
	_ bin.BareEncoder = &GetMessageEmbeddingCodeRequest{}
	_ bin.BareDecoder = &GetMessageEmbeddingCodeRequest{}
)

func (g *GetMessageEmbeddingCodeRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.ForAlbum == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageEmbeddingCodeRequest) String() string {
	if g == nil {
		return "GetMessageEmbeddingCodeRequest(nil)"
	}
	type Alias GetMessageEmbeddingCodeRequest
	return fmt.Sprintf("GetMessageEmbeddingCodeRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageEmbeddingCodeRequest) TypeID() uint32 {
	return GetMessageEmbeddingCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageEmbeddingCodeRequest) TypeName() string {
	return "getMessageEmbeddingCode"
}

// TypeInfo returns info about TL type.
func (g *GetMessageEmbeddingCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageEmbeddingCode",
		ID:   GetMessageEmbeddingCodeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ForAlbum",
			SchemaName: "for_album",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageEmbeddingCodeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageEmbeddingCode#62a4cd09 as nil")
	}
	b.PutID(GetMessageEmbeddingCodeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageEmbeddingCodeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageEmbeddingCode#62a4cd09 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutLong(g.MessageID)
	b.PutBool(g.ForAlbum)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageEmbeddingCodeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageEmbeddingCode#62a4cd09 to nil")
	}
	if err := b.ConsumeID(GetMessageEmbeddingCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageEmbeddingCodeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageEmbeddingCode#62a4cd09 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: field for_album: %w", err)
		}
		g.ForAlbum = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageEmbeddingCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageEmbeddingCode#62a4cd09 as nil")
	}
	b.ObjStart()
	b.PutID("getMessageEmbeddingCode")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_id")
	b.PutLong(g.MessageID)
	b.FieldStart("for_album")
	b.PutBool(g.ForAlbum)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageEmbeddingCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageEmbeddingCode#62a4cd09 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageEmbeddingCode"); err != nil {
				return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: field message_id: %w", err)
			}
			g.MessageID = value
		case "for_album":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageEmbeddingCode#62a4cd09: field for_album: %w", err)
			}
			g.ForAlbum = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageEmbeddingCodeRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageEmbeddingCodeRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetForAlbum returns value of ForAlbum field.
func (g *GetMessageEmbeddingCodeRequest) GetForAlbum() (value bool) {
	return g.ForAlbum
}

// GetMessageEmbeddingCode invokes method getMessageEmbeddingCode#62a4cd09 returning error if any.
func (c *Client) GetMessageEmbeddingCode(ctx context.Context, request *GetMessageEmbeddingCodeRequest) (*Text, error) {
	var result Text

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
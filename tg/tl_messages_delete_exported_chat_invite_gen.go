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

// MessagesDeleteExportedChatInviteRequest represents TL type `messages.deleteExportedChatInvite#d464a42b`.
//
// See https://core.telegram.org/method/messages.deleteExportedChatInvite for reference.
type MessagesDeleteExportedChatInviteRequest struct {
	// Peer field of MessagesDeleteExportedChatInviteRequest.
	Peer InputPeerClass
	// Link field of MessagesDeleteExportedChatInviteRequest.
	Link string
}

// MessagesDeleteExportedChatInviteRequestTypeID is TL type id of MessagesDeleteExportedChatInviteRequest.
const MessagesDeleteExportedChatInviteRequestTypeID = 0xd464a42b

// Ensuring interfaces in compile-time for MessagesDeleteExportedChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesDeleteExportedChatInviteRequest{}
	_ bin.Decoder     = &MessagesDeleteExportedChatInviteRequest{}
	_ bin.BareEncoder = &MessagesDeleteExportedChatInviteRequest{}
	_ bin.BareDecoder = &MessagesDeleteExportedChatInviteRequest{}
)

func (d *MessagesDeleteExportedChatInviteRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteExportedChatInviteRequest) String() string {
	if d == nil {
		return "MessagesDeleteExportedChatInviteRequest(nil)"
	}
	type Alias MessagesDeleteExportedChatInviteRequest
	return fmt.Sprintf("MessagesDeleteExportedChatInviteRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteExportedChatInviteRequest from given interface.
func (d *MessagesDeleteExportedChatInviteRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetLink() (value string)
}) {
	d.Peer = from.GetPeer()
	d.Link = from.GetLink()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteExportedChatInviteRequest) TypeID() uint32 {
	return MessagesDeleteExportedChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteExportedChatInviteRequest) TypeName() string {
	return "messages.deleteExportedChatInvite"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteExportedChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteExportedChatInvite",
		ID:   MessagesDeleteExportedChatInviteRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteExportedChatInviteRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteExportedChatInvite#d464a42b as nil")
	}
	b.PutID(MessagesDeleteExportedChatInviteRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteExportedChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteExportedChatInvite#d464a42b as nil")
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode messages.deleteExportedChatInvite#d464a42b: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteExportedChatInvite#d464a42b: field peer: %w", err)
	}
	b.PutString(d.Link)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteExportedChatInviteRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteExportedChatInvite#d464a42b to nil")
	}
	if err := b.ConsumeID(MessagesDeleteExportedChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteExportedChatInvite#d464a42b: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteExportedChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteExportedChatInvite#d464a42b to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteExportedChatInvite#d464a42b: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteExportedChatInvite#d464a42b: field link: %w", err)
		}
		d.Link = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *MessagesDeleteExportedChatInviteRequest) GetPeer() (value InputPeerClass) {
	return d.Peer
}

// GetLink returns value of Link field.
func (d *MessagesDeleteExportedChatInviteRequest) GetLink() (value string) {
	return d.Link
}

// MessagesDeleteExportedChatInvite invokes method messages.deleteExportedChatInvite#d464a42b returning error if any.
//
// See https://core.telegram.org/method/messages.deleteExportedChatInvite for reference.
func (c *Client) MessagesDeleteExportedChatInvite(ctx context.Context, request *MessagesDeleteExportedChatInviteRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}

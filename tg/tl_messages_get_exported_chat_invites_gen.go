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

// MessagesGetExportedChatInvitesRequest represents TL type `messages.getExportedChatInvites#a2b5a3f6`.
//
// See https://core.telegram.org/method/messages.getExportedChatInvites for reference.
type MessagesGetExportedChatInvitesRequest struct {
	// Flags field of MessagesGetExportedChatInvitesRequest.
	Flags bin.Fields
	// Revoked field of MessagesGetExportedChatInvitesRequest.
	Revoked bool
	// Peer field of MessagesGetExportedChatInvitesRequest.
	Peer InputPeerClass
	// AdminID field of MessagesGetExportedChatInvitesRequest.
	AdminID InputUserClass
	// OffsetDate field of MessagesGetExportedChatInvitesRequest.
	//
	// Use SetOffsetDate and GetOffsetDate helpers.
	OffsetDate int
	// OffsetLink field of MessagesGetExportedChatInvitesRequest.
	//
	// Use SetOffsetLink and GetOffsetLink helpers.
	OffsetLink string
	// Limit field of MessagesGetExportedChatInvitesRequest.
	Limit int
}

// MessagesGetExportedChatInvitesRequestTypeID is TL type id of MessagesGetExportedChatInvitesRequest.
const MessagesGetExportedChatInvitesRequestTypeID = 0xa2b5a3f6

// Ensuring interfaces in compile-time for MessagesGetExportedChatInvitesRequest.
var (
	_ bin.Encoder     = &MessagesGetExportedChatInvitesRequest{}
	_ bin.Decoder     = &MessagesGetExportedChatInvitesRequest{}
	_ bin.BareEncoder = &MessagesGetExportedChatInvitesRequest{}
	_ bin.BareDecoder = &MessagesGetExportedChatInvitesRequest{}
)

func (g *MessagesGetExportedChatInvitesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Revoked == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.AdminID == nil) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.OffsetLink == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetExportedChatInvitesRequest) String() string {
	if g == nil {
		return "MessagesGetExportedChatInvitesRequest(nil)"
	}
	type Alias MessagesGetExportedChatInvitesRequest
	return fmt.Sprintf("MessagesGetExportedChatInvitesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetExportedChatInvitesRequest from given interface.
func (g *MessagesGetExportedChatInvitesRequest) FillFrom(from interface {
	GetRevoked() (value bool)
	GetPeer() (value InputPeerClass)
	GetAdminID() (value InputUserClass)
	GetOffsetDate() (value int, ok bool)
	GetOffsetLink() (value string, ok bool)
	GetLimit() (value int)
}) {
	g.Revoked = from.GetRevoked()
	g.Peer = from.GetPeer()
	g.AdminID = from.GetAdminID()
	if val, ok := from.GetOffsetDate(); ok {
		g.OffsetDate = val
	}

	if val, ok := from.GetOffsetLink(); ok {
		g.OffsetLink = val
	}

	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetExportedChatInvitesRequest) TypeID() uint32 {
	return MessagesGetExportedChatInvitesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetExportedChatInvitesRequest) TypeName() string {
	return "messages.getExportedChatInvites"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetExportedChatInvitesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getExportedChatInvites",
		ID:   MessagesGetExportedChatInvitesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoked",
			SchemaName: "revoked",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "OffsetLink",
			SchemaName: "offset_link",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetExportedChatInvitesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getExportedChatInvites#a2b5a3f6 as nil")
	}
	b.PutID(MessagesGetExportedChatInvitesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetExportedChatInvitesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getExportedChatInvites#a2b5a3f6 as nil")
	}
	if !(g.Revoked == false) {
		g.Flags.Set(3)
	}
	if !(g.OffsetDate == 0) {
		g.Flags.Set(2)
	}
	if !(g.OffsetLink == "") {
		g.Flags.Set(2)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvites#a2b5a3f6: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvites#a2b5a3f6: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvites#a2b5a3f6: field peer: %w", err)
	}
	if g.AdminID == nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvites#a2b5a3f6: field admin_id is nil")
	}
	if err := g.AdminID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvites#a2b5a3f6: field admin_id: %w", err)
	}
	if g.Flags.Has(2) {
		b.PutInt(g.OffsetDate)
	}
	if g.Flags.Has(2) {
		b.PutString(g.OffsetLink)
	}
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetExportedChatInvitesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getExportedChatInvites#a2b5a3f6 to nil")
	}
	if err := b.ConsumeID(MessagesGetExportedChatInvitesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetExportedChatInvitesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getExportedChatInvites#a2b5a3f6 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: field flags: %w", err)
		}
	}
	g.Revoked = g.Flags.Has(3)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: field admin_id: %w", err)
		}
		g.AdminID = value
	}
	if g.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	if g.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: field offset_link: %w", err)
		}
		g.OffsetLink = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvites#a2b5a3f6: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetRevoked sets value of Revoked conditional field.
func (g *MessagesGetExportedChatInvitesRequest) SetRevoked(value bool) {
	if value {
		g.Flags.Set(3)
		g.Revoked = true
	} else {
		g.Flags.Unset(3)
		g.Revoked = false
	}
}

// GetRevoked returns value of Revoked conditional field.
func (g *MessagesGetExportedChatInvitesRequest) GetRevoked() (value bool) {
	return g.Flags.Has(3)
}

// GetPeer returns value of Peer field.
func (g *MessagesGetExportedChatInvitesRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetAdminID returns value of AdminID field.
func (g *MessagesGetExportedChatInvitesRequest) GetAdminID() (value InputUserClass) {
	return g.AdminID
}

// SetOffsetDate sets value of OffsetDate conditional field.
func (g *MessagesGetExportedChatInvitesRequest) SetOffsetDate(value int) {
	g.Flags.Set(2)
	g.OffsetDate = value
}

// GetOffsetDate returns value of OffsetDate conditional field and
// boolean which is true if field was set.
func (g *MessagesGetExportedChatInvitesRequest) GetOffsetDate() (value int, ok bool) {
	if !g.Flags.Has(2) {
		return value, false
	}
	return g.OffsetDate, true
}

// SetOffsetLink sets value of OffsetLink conditional field.
func (g *MessagesGetExportedChatInvitesRequest) SetOffsetLink(value string) {
	g.Flags.Set(2)
	g.OffsetLink = value
}

// GetOffsetLink returns value of OffsetLink conditional field and
// boolean which is true if field was set.
func (g *MessagesGetExportedChatInvitesRequest) GetOffsetLink() (value string, ok bool) {
	if !g.Flags.Has(2) {
		return value, false
	}
	return g.OffsetLink, true
}

// GetLimit returns value of Limit field.
func (g *MessagesGetExportedChatInvitesRequest) GetLimit() (value int) {
	return g.Limit
}

// MessagesGetExportedChatInvites invokes method messages.getExportedChatInvites#a2b5a3f6 returning error if any.
//
// See https://core.telegram.org/method/messages.getExportedChatInvites for reference.
func (c *Client) MessagesGetExportedChatInvites(ctx context.Context, request *MessagesGetExportedChatInvitesRequest) (*MessagesExportedChatInvites, error) {
	var result MessagesExportedChatInvites

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

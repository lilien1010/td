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

// MessagesGetSuggestedDialogFiltersRequest represents TL type `messages.getSuggestedDialogFilters#a29cd42c`.
// Get suggested folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.getSuggestedDialogFilters for reference.
type MessagesGetSuggestedDialogFiltersRequest struct {
}

// MessagesGetSuggestedDialogFiltersRequestTypeID is TL type id of MessagesGetSuggestedDialogFiltersRequest.
const MessagesGetSuggestedDialogFiltersRequestTypeID = 0xa29cd42c

// Ensuring interfaces in compile-time for MessagesGetSuggestedDialogFiltersRequest.
var (
	_ bin.Encoder     = &MessagesGetSuggestedDialogFiltersRequest{}
	_ bin.Decoder     = &MessagesGetSuggestedDialogFiltersRequest{}
	_ bin.BareEncoder = &MessagesGetSuggestedDialogFiltersRequest{}
	_ bin.BareDecoder = &MessagesGetSuggestedDialogFiltersRequest{}
)

func (g *MessagesGetSuggestedDialogFiltersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSuggestedDialogFiltersRequest) String() string {
	if g == nil {
		return "MessagesGetSuggestedDialogFiltersRequest(nil)"
	}
	type Alias MessagesGetSuggestedDialogFiltersRequest
	return fmt.Sprintf("MessagesGetSuggestedDialogFiltersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetSuggestedDialogFiltersRequest) TypeID() uint32 {
	return MessagesGetSuggestedDialogFiltersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetSuggestedDialogFiltersRequest) TypeName() string {
	return "messages.getSuggestedDialogFilters"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetSuggestedDialogFiltersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getSuggestedDialogFilters",
		ID:   MessagesGetSuggestedDialogFiltersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetSuggestedDialogFiltersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSuggestedDialogFilters#a29cd42c as nil")
	}
	b.PutID(MessagesGetSuggestedDialogFiltersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetSuggestedDialogFiltersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSuggestedDialogFilters#a29cd42c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSuggestedDialogFiltersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSuggestedDialogFilters#a29cd42c to nil")
	}
	if err := b.ConsumeID(MessagesGetSuggestedDialogFiltersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSuggestedDialogFilters#a29cd42c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetSuggestedDialogFiltersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSuggestedDialogFilters#a29cd42c to nil")
	}
	return nil
}

// MessagesGetSuggestedDialogFilters invokes method messages.getSuggestedDialogFilters#a29cd42c returning error if any.
// Get suggested folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.getSuggestedDialogFilters for reference.
func (c *Client) MessagesGetSuggestedDialogFilters(ctx context.Context) ([]DialogFilterSuggested, error) {
	var result DialogFilterSuggestedVector

	request := &MessagesGetSuggestedDialogFiltersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []DialogFilterSuggested(result.Elems), nil
}
